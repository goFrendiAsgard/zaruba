package linescmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var printCmd = &cobra.Command{
	Use:     "print <lines> [fileName]",
	Short:   "Print lines",
	Aliases: []string{"write"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonString := args[0]
		fileName := ""
		if len(args) > 1 {
			fileName = args[1]
		}
		util := dsl.NewDSLUtil()
		if err := util.File.WriteLines(fileName, jsonString, 0755); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
