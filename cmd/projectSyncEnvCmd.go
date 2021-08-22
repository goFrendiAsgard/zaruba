package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var projectSyncEnvCmd = &cobra.Command{
	Use:   "syncEnv <projectFile>",
	Short: "Update every 'run*' task's environment",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		project, err := getProject(decoration, projectFile)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		if err = project.Init(); err != nil {
			exit(cmd, logger, decoration, err)
		}
		if err = config.SyncProjectEnv(project); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
