package taskcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getIconCmd = &cobra.Command{
	Use:   "getIcon <taskName> [projectFile]",
	Short: "get task icon",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		taskName := args[0]
		projectFilePath, err := cmdHelper.GetProjectRelFilePath(args, 1, "index.zaruba.yaml", "index.zaruba.yml")
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		util := dsl.NewDSLUtil()
		icon, err := util.Project.Task.GetIcon(taskName, projectFilePath)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(icon)
	},
}
