package taskcmd

import (
	"path/filepath"

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
		projectFile := "index.zaruba.yaml"
		if len(args) > 2 {
			projectFile = args[2]
		}
		projectFile, err = filepath.Abs(projectFile)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if err = util.Project.Task.AddParents(taskName, parentList, projectFile); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
