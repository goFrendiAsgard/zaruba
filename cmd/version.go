package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/logger"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Zaruba",
	Long:  `All software has versions. This is Zaruba's`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Zaruba v0.0.0 -- [prototype]")
	},
}
