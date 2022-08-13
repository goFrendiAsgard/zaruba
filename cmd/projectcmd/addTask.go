package projectcmd

import (
	"path/filepath"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var addTaskCmd = &cobra.Command{
	Use:   "addTask <taskName> [taskFile]",
	Short: "Add task to current project",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		taskName := args[0]
		taskFilePath := "index.zaruba.yaml"
		if len(args) > 1 {
			taskFilePath = args[1]
		}
		var err error
		taskFilePath, err = filepath.Abs(taskFilePath)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		util := dsl.NewDSLUtil()
		if err = util.Project.AddTaskIfNotExist(taskName, taskFilePath); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
