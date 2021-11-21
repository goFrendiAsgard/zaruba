package projectcmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var setValueCmd = &cobra.Command{
	Use:   "setValue <valueFile> <key> <value>",
	Short: "Set project value",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		util := core.NewCoreUtil()
		if err := util.Project.SetValue(args[0], args[1], args[2]); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
