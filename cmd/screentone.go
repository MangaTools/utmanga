package cmd

import (
	"errors"
	"os"
	"runtime"

	"github.com/MangaTools/utmanga/dto"
	"github.com/MangaTools/utmanga/dto/validator"
	"github.com/MangaTools/utmanga/executor"
	"github.com/MangaTools/utmanga/screentone"
	"github.com/spf13/cobra"
)

var (
	screentoneCmd = &cobra.Command{
		Use:   "screentone",
		Short: "This command starts the process of overlaying a screentone on the images in the specified folder.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return errors.Join(
				validator.ValidateExecutorSettings(screentoneExecutorSettings),
				validator.ValidateScreentone(screentoneSettings))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			screentoner := screentone.NewScreentoner(screentoneSettings)
			runner := executor.NewExecutor(screentoneExecutorSettings, screentoner.Execute)

			return runner.Execute()
		},
	}

	screentonePipeCmd = &cobra.Command{
		Use:   "pipe",
		Short: "This command starts processing of overlaying a screentone of a single file from the stdin stream and outputs to stdout stream.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validator.ValidateScreentone(screentoneSettings)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			screentoner := screentone.NewScreentoner(screentoneSettings)

			return executor.ExecutePipe(screentoner.Execute, executor.DefaultEncoder, os.Stdin, os.Stdout)
		},
	}
)

var (
	screentoneExecutorSettings dto.ExecutorSettings
	screentoneSettings         dto.ScreentoneSettings
)

func init() {
	screentoneCmd.AddCommand(screentonePipeCmd)

	screentoneCmd.PersistentFlags().IntVarP(&screentoneSettings.DotSize, "dot_size", "d", 6,
		"Max dot size in pixels. Value must be in range [2;100]. This flag is responsible for how big your screentone will be.")
	screentoneCmd.MarkFlagRequired("dot_size")

	screentoneCmd.Flags().IntVarP(&screentoneExecutorSettings.MaxGoroutines, "threads", "t", runtime.NumCPU(),
		"Number of simultaneously processed images. The default value is equal to the number of logical processor cores.")

	screentoneCmd.Flags().StringVarP(&screentoneExecutorSettings.InputPath, "input", "i", "",
		"The path to the folder with the images. Required.")
	screentoneCmd.MarkFlagRequired("input")

	screentoneCmd.Flags().BoolVarP(&screentoneExecutorSettings.Recursive, "recursive", "r", false,
		"Recursive image search.")
	screentoneCmd.Flags().StringVarP(&screentoneExecutorSettings.OutputPath, "out", "o", "",
		`Path to the folder with final images (if the folder does not exist, it will be created).
	If not specified, it is written to the input folder. The output will be .png files. If there are already files with the same name in the folder, the files are overwritten.`)
}
