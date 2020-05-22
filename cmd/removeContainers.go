package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
	"github.com/state-alchemists/zaruba/modules/runner"
)

func init() {
	rootCmd.AddCommand(removeContainersCmd)
}

var removeContainersCmd = &cobra.Command{
	Use:   "remove-containers [project-dir]",
	Short: "Remove all containers.",
	Long:  "Remove all containers defined in a project.",
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
		p, err := config.CreateProjectConfig(projectDir)
		if err != nil {
			logger.Fatal(err)
		}
		// invoke action
		if err = runner.RemoveContainers(projectDir, p); err != nil {
			logger.Fatal(err)
		}
	},
}
