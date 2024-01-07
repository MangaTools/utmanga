package cmd

import (
	"runtime"

	"github.com/MangaTools/utmanga/dto"
	"github.com/spf13/cobra"
)

func addExecutorFlagsToCommand(command *cobra.Command, settings *dto.ExecutorSettings) {
	command.Flags().IntVarP(&settings.MaxGoroutines, "threads", "t", runtime.NumCPU(),
		"Number of simultaneously processed images. The default value is equal to the number of logical processor cores.")

	command.Flags().StringVarP(&settings.InputPath, "input", "i", "",
		"The path to the folder with the images. Required.")
	command.MarkFlagRequired("input")

	command.Flags().BoolVarP(&settings.Recursive, "recursive", "r", false,
		"Recursive image search.")
	command.Flags().StringVarP(&settings.OutputPath, "out", "o", "",
		`Path to the folder with final images (if the folder does not exist, it will be created).
	If not specified, it is written to the input folder. If there are already files with the same name in the folder, the files are overwritten.`)
}
