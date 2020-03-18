package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/logger"
	"github.com/state-alchemists/zaruba/modules/puller"
)

func init() {
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use:   "pull [project-dir]",
	Short: "Pull from subtrees",
	Long:  `Zaruba will pull from subtrees`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir := "."
		if len(args) > 0 {
			projectDir = args[0]
		}
		// invoke action
		if err := puller.Pull(projectDir); err != nil {
			logger.Fatal(err)
		}
	},
}
