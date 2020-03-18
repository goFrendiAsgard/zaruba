package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/initiator"
	"github.com/state-alchemists/zaruba/modules/logger"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [project-dir]",
	Short: "Init a project",
	Long:  `Zaruba will init a project`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir := "."
		if len(args) > 0 {
			projectDir = args[0]
		}
		// invoke action
		if err := initiator.Init(projectDir); err != nil {
			logger.Fatal(err)
		}
	},
}
