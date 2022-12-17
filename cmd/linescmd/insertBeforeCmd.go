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
	Use:   "insertBefore <jsonStrList> <jsonStrListNewLines>",
	Short: "Insert newLine before lines[index]",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListNewLines := args[0], args[1]
		newJsonStrList, err := util.Json.List.InsertLineBeforeIndex(jsonStrList, *insertBeforeIndex, jsonStrListNewLines)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newJsonStrList)
	},
}
