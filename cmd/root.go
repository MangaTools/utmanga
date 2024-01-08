package cmd

import (
	"github.com/MangaTools/utmanga/globals"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "utmanga",
	Short:   "utmanga is a cli application to sum up tools for bw manga image processing. For available command enter 'utmanga --help'",
	Version: globals.AppVersion,
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(screentoneCmd, resizeCmd, levelCmd, pipelineCmd)
}

func Execute() {
	rootCmd.Execute()
}
