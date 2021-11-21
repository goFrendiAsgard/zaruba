package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var taskSyncEnvCmd = &cobra.Command{
	Use:   "syncEnv <projectFile> <taskName>",
	Short: "Update task's environment",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		taskName := args[1]
		util := core.NewCoreUtil()
		if err = util.Project.Task.Env.Sync(projectFile, taskName); err != nil {
			exit(cmd, args, logger, decoration, err)
		}
	},
}
