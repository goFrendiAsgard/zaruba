package projectcmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var addTaskIfNotExistCmd = &cobra.Command{
	Use:   "addTaskIfNotExist <taskFilePath> <taskName>",
	Short: "Add task to project",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		taskFilePath, err := filepath.Abs(args[0])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		taskName := args[1]
		util := core.NewCoreUtil()
		if err = util.Project.AddTaskIfNotExist(taskFilePath, taskName); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
