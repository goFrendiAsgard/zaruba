package projectcmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var setValueCmd = &cobra.Command{
	Use:   "setValue <key> <value> [projectFile]",
	Short: "Set project value",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		key := args[0]
		value := args[1]
		projectFilePath, err := cmdHelper.GetProjectRelFilePath(args, 2, "index.zaruba.yaml", "index.zaruba.yml")
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		util := dsl.NewDSLUtil()
		if err := util.Project.SetValue(key, value, projectFilePath); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
