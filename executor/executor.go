package executor

import (
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"

	"github.com/MangaTools/utmanga/dto"
	"github.com/schollz/progressbar/v3"
	"github.com/sourcegraph/conc/pool"
)

type ExecutionFunc func(image image.Image) (image.Image, error)

type Executor struct {
	settings dto.ExecutorSettings

	execute ExecutionFunc
	encode  EncoderFunc
}

// NewExecutor creates Executor with settings and execution func. All args should be valid!
func NewExecutor(settings dto.ExecutorSettings, executionFunc ExecutionFunc) *Executor {
	return &Executor{
		settings: settings,
		execute:  executionFunc,
		encode:   DefaultEncoder,
	}
}

func (e *Executor) Execute() error {
	inputFolder, err := filepath.Abs(e.settings.InputFolder)
	if err != nil {
		return fmt.Errorf("get absolute path of input folder: %w", err)
	}

	images, err := getImagesPaths(inputFolder, acceptedImageExtensions, e.settings.Recursive)
	if err != nil {
		return err
	}

	outputFolder, err := filepath.Abs(e.settings.OutputFolder)
	if err != nil {
		return fmt.Errorf("get absolute path of output folder: %w", err)
	}

	items := len(images)

	wgPool := pool.New().WithMaxGoroutines(e.settings.MaxGoroutines)
	bar := progressbar.Default(int64(items), "Executing...")

	for _, imageFile := range images {
		imageInputPath := filepath.Join(inputFolder, imageFile)
		imageOutputPath := filepath.Join(outputFolder, imageFile)

		err := os.MkdirAll(filepath.Dir(imageOutputPath), 0o755)
		if err != nil {
			return fmt.Errorf("create output dir \"%s\": %w", filepath.Dir(imageOutputPath), err)
		}

		wgPool.Go(func() {
			err = e.executeFile(imageInputPath, imageOutputPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "can not handle image \"%s\": %s", imageInputPath, err.Error())
			}

			bar.Add(1)
		})
	}

	wgPool.Wait()

	return nil
}

func (e *Executor) executeFile(inputFilePath, outputFilePath string) error {
	inFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("open input file %s to read: %w", inputFilePath, err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("create or truncate output file %s: %w", outputFilePath, err)
	}
	defer outFile.Close()

	err = ExecutePipe(e.execute, e.encode, inFile, outFile)
	if err != nil {
		return fmt.Errorf("execute file %s to %s: %w", inputFilePath, outputFilePath, err)
	}

	return nil
}

func ExecutePipe(execute ExecutionFunc, encode EncoderFunc, input io.Reader, output io.Writer) error {
	img, err := readImageFromReader(input)
	if err != nil {
		return fmt.Errorf("read input to image struct: %w", err)
	}

	resultImg, err := execute(img)
	if err != nil {
		return fmt.Errorf("execute func on image: %w", err)
	}

	err = encode(output, resultImg)
	if err != nil {
		return fmt.Errorf("encode result image to output: %w", err)
	}

	return nil
}
