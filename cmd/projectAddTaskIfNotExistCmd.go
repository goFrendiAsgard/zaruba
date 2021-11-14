package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var projectAddTaskIfNotExistCmd = &cobra.Command{
	Use:   "addTaskIfNotExist <taskFilePath> <taskName>",
	Short: "Add task to project",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		taskFilePath, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		taskName := args[1]
		if err = core.AddTaskIfNotExist(taskFilePath, taskName); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
