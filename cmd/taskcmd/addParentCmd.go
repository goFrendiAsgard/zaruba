package taskcmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var addParentCmd = &cobra.Command{
	Use:   "addParent <projectFile> <taskName> <newParentNames>",
	Short: "Add task parent",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		taskName := args[1]
		util := core.NewCoreUtil()
		parentNames, err := util.Json.List.GetStringList(args[2])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if err = util.Project.Task.AddParents(projectFile, taskName, parentNames); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
