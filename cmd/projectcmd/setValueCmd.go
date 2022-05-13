package projectcmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var setValueCmd = &cobra.Command{
	Use:   "setValue <key> <value> [valueFile]",
	Short: "Set project value",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		key := args[0]
		value := args[1]
		valueFile := "default.value.yaml"
		if len(args) > 2 {
			valueFile = args[2]
		}
		util := core.NewCoreUtil()
		if err := util.Project.SetValue(key, value, valueFile); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
