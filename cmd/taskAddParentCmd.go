package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var taskAddParentCmd = &cobra.Command{
	Use:   "addParent <projectFile> <taskName> <newParentNames>",
	Short: "Add task parent",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		taskName := args[1]
		util := core.NewCoreUtil()
		parentNames, err := util.Json.List.GetStringList(args[2])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		if err = util.Project.Task.AddParents(projectFile, taskName, parentNames); err != nil {
			exit(cmd, args, logger, decoration, err)
		}
	},
}
