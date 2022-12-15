package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertAfterIndex *int
var insertAfterCmd = &cobra.Command{
	Use:   "insertAfter <jsonStrList> <jsonStrListNewLines>",
	Short: "Insert newLine after lines[index]",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListNewLines := args[0], args[1]
		newJsonStrList, err := util.Json.List.InsertLineAfterIndex(jsonStrList, *insertAfterIndex, jsonStrListNewLines)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newJsonStrList)
	},
}
