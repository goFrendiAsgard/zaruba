package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var projectSyncEnvCmd = &cobra.Command{
	Use:   "syncEnv <projectFile>",
	Short: "Update every task's environment",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		util := core.NewCoreUtil()
		if err = util.Project.SyncTasksEnv(projectFile); err != nil {
			exit(cmd, args, logger, decoration, err)
		}
	},
}
