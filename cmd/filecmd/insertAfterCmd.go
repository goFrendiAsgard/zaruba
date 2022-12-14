package filecmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertAfterIndex *int
var insertAfterCmd = &cobra.Command{
	Use:   "insertAfter <fileName> <newContent>",
	Short: "Insert newContent to fileName after index",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonLines, newContent := args[0], args[1]
		if err := util.File.InsertLineAfterIndex(jsonLines, newContent, *insertAfterIndex); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
