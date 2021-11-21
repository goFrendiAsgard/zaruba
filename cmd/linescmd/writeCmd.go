package linescmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var writeCmd = &cobra.Command{
	Use:   "write <fileName> <jsonList>",
	Short: "Write list to file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		fileName, jsonString := args[0], args[1]
		util := core.NewCoreUtil()
		if err := util.File.WriteLines(fileName, jsonString, 0755); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
