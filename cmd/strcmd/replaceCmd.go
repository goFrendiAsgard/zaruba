package strcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var replaceCmd = &cobra.Command{
	Use:   "replace <string> [jsonMapReplacement]",
	Short: "Replace string by jsonMapReplacement",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		jsonMapReplacement := "{}"
		if len(args) > 1 {
			jsonMapReplacement = args[1]
		}
		util := dsl.NewDSLUtil()
		result, err := util.Json.Map.Replace(text, jsonMapReplacement)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(result)
	},
}
