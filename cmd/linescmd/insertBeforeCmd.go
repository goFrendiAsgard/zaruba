package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertBeforeIndex *int
var insertBeforeCmd = &cobra.Command{
	Use:   "insertBefore <lines> <newLine>",
	Short: "Insert newLine before lines[index]",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonLines, jsonReplacements := args[0], args[1]
		newJsonLines, err := util.Json.List.InsertLineBeforeIndex(jsonLines, *insertBeforeIndex, jsonReplacements)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newJsonLines)
	},
}
