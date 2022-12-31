package taskcmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	common "github.com/state-alchemists/zaruba/jsonutil/helper"
	"github.com/state-alchemists/zaruba/output"
)

var addParentsCmd = &cobra.Command{
	Use:     "addParents <taskName> {<jsonListParentTask> | <parentTaskName>} [projectFile]",
	Short:   "Add task parent",
	Aliases: []string{"addParent"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		taskName := args[0]
		util := dsl.NewDSLUtil()
		parentList, err := util.Json.ToStringList(args[1])
		if err != nil {
			parentList = common.StringList{args[1]}
		}
		projectFilePath, err := cmdHelper.GetProjectRelFilePath(args, 2, "index.zaruba.yaml", "index.zaruba.yml")
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		if err = util.Project.Task.AddParents(taskName, parentList, projectFilePath); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
