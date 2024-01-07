package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/MangaTools/utmanga/dto"
	"github.com/MangaTools/utmanga/dto/validator"
	"github.com/MangaTools/utmanga/executor"
	"github.com/MangaTools/utmanga/resize"
	"github.com/spf13/cobra"
)

var (
	resizeCmd = &cobra.Command{
		Use:   "resize",
		Short: "This command starts the process of resizing.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return errors.Join(validator.ValidateExecutorSettings(resizeExecutorSettings),
				validator.ValidateResizeSettings(resizeSettings))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			resizer := resize.NewResizer(resizeSettings)
			executor := executor.NewExecutor(resizeExecutorSettings, resizer.Execute)

			return executor.Execute()
		},
	}

	resizePipeCmd = &cobra.Command{
		Use:   "pipe",
		Short: "This command starts the process of resizing a single file from the stdin stream and outputs to stdout stream.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validator.ValidateResizeSettings(resizeSettings)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			resizer := resize.NewResizer(resizeSettings)

			return executor.ExecutePipe(resizer.Execute, executor.DefaultEncoder, os.Stdin, os.Stdout)
		},
	}
)

var (
	resizeExecutorSettings dto.ExecutorSettings
	resizeSettings         dto.ResizeSettings
)

func init() {
	resizeCmd.AddCommand(resizePipeCmd)

	addExecutorFlagsToCommand(resizeCmd, &resizeExecutorSettings)

	resizeCmd.PersistentFlags().StringVarP(&resizeSettings.Mode, "mode", "m", "",
		fmt.Sprintf("Mode of resize. Supported values: %s", strings.Join(dto.ResizeModeList, ", ")))
	resizeCmd.MarkFlagRequired("mode")

	resizeCmd.PersistentFlags().IntVarP(&resizeSettings.Steps, "steps", "s", 1, "How many steps resize execute to get to target size. Max is 10")

	resizeCmd.PersistentFlags().IntVarP(&resizeSettings.Coefficient, "coefficient", "c", 0, "Multiply or divide source size of image by x. Negative number is divide. Supported values: -4, -2 0, 2, 4")

	resizeCmd.PersistentFlags().IntVar(&resizeSettings.TargetHeigth, "height", 0, "Set target image height to x and save aspect ratio.")
	resizeCmd.PersistentFlags().IntVarP(&resizeSettings.TargetWidth, "width", "w", 0, "Set target image width to x and save aspect ratio.")

	resizeCmd.MarkFlagsOneRequired("coefficient", "height", "width")
	resizeCmd.MarkFlagsMutuallyExclusive("coefficient", "height", "width")
}
