package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var taskSyncEnvCmd = &cobra.Command{
	Use:   "syncEnv <projectFile> <taskName>",
	Short: "Update task's environment",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
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
		task, taskExist := project.Tasks[args[1]]
		if !taskExist {
			exit(cmd, logger, decoration, fmt.Errorf("task %s does not exist", args[1]))
		}
		if err = core.SyncTaskEnv(task); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
