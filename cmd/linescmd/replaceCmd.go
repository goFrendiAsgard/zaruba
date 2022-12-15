package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var replaceIndex *int
var replaceCmd = &cobra.Command{
	Use:   "replace <jsonStrList> <jsonStrListNewContent>",
	Short: "Replace lines[index] with replacements",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListNewContent := args[0], args[1]
		newJsonStrList, err := util.Json.List.ReplaceLineAtIndex(jsonStrList, *replaceIndex, jsonStrListNewContent)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newJsonStrList)
	},
}
