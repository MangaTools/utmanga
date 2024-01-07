package cmd

import (
	"errors"
	"os"

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
				validator.ValidateScreentoneSettings(screentoneSettings))
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
			return validator.ValidateScreentoneSettings(screentoneSettings)
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

	addExecutorFlagsToCommand(screentoneCmd, &screentoneExecutorSettings)
}
