package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
	"github.com/state-alchemists/zaruba/modules/runner"
)

func init() {
	rootCmd.AddCommand(stopContainersCmd)
}

var stopContainersCmd = &cobra.Command{
	Use:   "stop-containers [project-dir]",
	Short: "Stop containers in a project",
	Long:  `Zaruba will stop all containers related to a project`,
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
		if err = runner.StopContainers(projectDir, p); err != nil {
			logger.Fatal(err)
		}
	},
}