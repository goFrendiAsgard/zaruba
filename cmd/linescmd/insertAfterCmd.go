package linescmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertAfterCmd = &cobra.Command{
	Use:   "insertAfter <jsonList> <index> <newLine>",
	Short: "Insert newLine after lines[index]",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		util := dsl.NewDSLUtil()
		jsonLines, jsonReplacements := args[0], args[2]
		index, err := strconv.Atoi(args[1])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		newJsonLines, err := util.Json.List.InsertLineAfterIndex(jsonLines, index, jsonReplacements)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newJsonLines)
	},
}
