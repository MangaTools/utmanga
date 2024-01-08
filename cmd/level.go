package cmd

import (
	"errors"
	"os"

	"github.com/MangaTools/utmanga/dto"
	"github.com/MangaTools/utmanga/dto/validator"
	"github.com/MangaTools/utmanga/executor"
	"github.com/MangaTools/utmanga/level"
	"github.com/spf13/cobra"
)

var (
	levelCmd = &cobra.Command{
		Use:   "level",
		Short: "This command set up levels in image.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return errors.Join(validator.ValidateExecutorSettings(levelExecutorSettings),
				validator.ValidateLevelSettings(levelSettings))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			leveler := level.NewLeveler(levelSettings)
			runner := executor.NewExecutor(levelExecutorSettings, leveler.Execute)

			return runner.Execute()
		},
	}

	levelPipeCmd = &cobra.Command{
		Use:   "pipe",
		Short: "This command set up levels in image in stdin and outputs in stdout.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validator.ValidateLevelSettings(levelSettings)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			leveler := level.NewLeveler(levelSettings)
			return executor.ExecutePipe(leveler.Execute, executor.DefaultEncoder, os.Stdin, os.Stdout)
		},
	}
)

var (
	levelExecutorSettings dto.ExecutorSettings
	levelSettings         dto.LevelSettings
)

func init() {
	levelCmd.AddCommand(levelPipeCmd)

	addExecutorFlagsToCommand(levelCmd, &levelExecutorSettings)

	levelCmd.PersistentFlags().IntVar(&levelSettings.LowInput, "low_input", 0, "low input")
	levelCmd.PersistentFlags().IntVar(&levelSettings.HighInput, "high_input", 255, "high input")
	levelCmd.PersistentFlags().Float64Var(&levelSettings.Gamma, "gamma", 1.0, "gamma")

	levelCmd.PersistentFlags().IntVar(&levelSettings.LowOutput, "low_output", 0, "low output")
	levelCmd.PersistentFlags().IntVar(&levelSettings.HighOutput, "high_output", 255, "high output")

	levelCmd.PersistentFlags().IntVar(&levelSettings.DiapasonBlack, "diapason_black", -1, "diapason black. -1 = off.")
	levelCmd.PersistentFlags().IntVar(&levelSettings.DiapasonWhite, "diapason_white", -1, "diapason white. -1 = off.")
}
