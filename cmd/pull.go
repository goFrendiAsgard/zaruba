package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/config"
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
		// get absolute project dir and project config
		projectDir, err := filepath.Abs(projectDir)
		if err != nil {
			logger.Fatal(err)
		}
		p, err := config.NewProjectConfig(projectDir)
		if err != nil {
			logger.Fatal(err)
		}
		// invoke action
		if err := puller.Pull(projectDir, p); err != nil {
			logger.Fatal(err)
		}
	},
}
