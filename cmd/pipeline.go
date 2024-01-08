package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MangaTools/utmanga/dto"
	"github.com/MangaTools/utmanga/dto/validator"
	"github.com/MangaTools/utmanga/executor"
	"github.com/MangaTools/utmanga/pipeline"
	"github.com/spf13/cobra"
)

var (
	pipelineCmd = &cobra.Command{
		Use:   "pipeline",
		Short: "This command runs execution steps according input json file on each image in input folder.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validator.ValidateExecutorSettings(pipelineExecutionSettings); err != nil {
				return err
			}

			return readAndValidateConfigFileFromFile()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			pipliner := pipeline.NewPipeliner(pipelineFileSettings)
			runner := executor.NewExecutor(pipelineExecutionSettings, pipliner.Execute)

			return runner.Execute()
		},
	}

	pipelineExampleCmd = &cobra.Command{
		Use:   "example",
		Short: "This command print example with all possible settings and types of pipeline json config in stdout.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(os.Stdout, generateExamplePipelineJson())
		},
	}

	pipelinePipeCmd = &cobra.Command{
		Use:   "pipe",
		Short: "This command runs execution steps according input json file on image in stdin and output result in stdout.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return readAndValidateConfigFileFromFile()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			pipliner := pipeline.NewPipeliner(pipelineFileSettings)

			return executor.ExecutePipe(pipliner.Execute, executor.DefaultEncoder, os.Stdin, os.Stdout)
		},
	}
)

var (
	pipelineExecutionSettings dto.ExecutorSettings
	pipelineFileSettings      dto.PipelineFileSettings

	pipelineConfigPath string
)

func init() {
	pipelineCmd.AddCommand(pipelineExampleCmd, pipelinePipeCmd)

	addExecutorFlagsToCommand(pipelineCmd, &pipelineExecutionSettings)
	addPipelineConfigPathCommand(pipelineCmd, pipelinePipeCmd)
}

func addPipelineConfigPathCommand(commands ...*cobra.Command) {
	for _, command := range commands {
		command.Flags().StringVarP(&pipelineConfigPath, "config_path", "c", "",
			"Path to json config file for pipeline. To check example config use 'utmanga pipeline example'. Steps are executed in original order of json array. For complete information about the parameter, please refer to the help tab of the corresponding command")
		command.MarkFlagRequired("config_path")
	}
}

func readAndValidateConfigFileFromFile() error {
	configFile, err := os.ReadFile(pipelineConfigPath)
	if err != nil {
		return fmt.Errorf("read pipeline config file: %w", err)
	}

	err = json.Unmarshal(configFile, &pipelineFileSettings)
	if err != nil {
		return fmt.Errorf("unmarshall pipeline config file: %w", err)
	}

	return validator.ValidatePipelineSettings(pipelineFileSettings)
}

func generateExamplePipelineJson() string {
	data, _ := json.MarshalIndent(examplePipelineFile, "", "    ")

	return string(data)
}

var (
	examplePipelineFile = dto.PipelineFileSettings{
		Items: []dto.PipelineItem{
			{
				Type:     dto.PipelineTypeLevel,
				Settings: levelExample,
			},
			{
				Type:     dto.PipelineTypeScrentone,
				Settings: screentoneExample,
			},
			{
				Type:     dto.PipelineTypeResize,
				Settings: resizeExample,
			},
		},
	}

	levelExample = dto.LevelSettings{
		LowInput:      0,
		HighInput:     255,
		Gamma:         1,
		LowOutput:     0,
		HighOutput:    255,
		DiapasonBlack: -1,
		DiapasonWhite: -1,
	}

	screentoneExample = dto.ScreentoneSettings{
		DotSize: 6,
	}

	resizeExample = dto.ResizeSettings{
		Mode:         dto.ResizeModeBox,
		Coefficient:  0,
		TargetHeigth: 2000,
		TargetWidth:  0,
		Steps:        1,
	}
)
