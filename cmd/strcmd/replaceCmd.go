package strcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var replaceExample = `
> zaruba str replace 'Capital of country is city' '{"country": "Indonesia", "city": "Jakarta"}'
Capital of Indonesia is Jakarta

> zaruba str replace 'Capital of country is city' country Japan city Tokyo
Capital of Japan is Tokyo
`

var replaceCmd = &cobra.Command{
	Use:     "replace <string> [{<jsonMapReplacement> | <key> <value>}]",
	Short:   "Replace string by jsonMapReplacement",
	Example: replaceExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		jsonMapReplacement := "{}"
		if len(args) > 1 {
			var err error
			jsonMapReplacement, err = cmdHelper.ArgToJsonReplacementMap(args, 1)
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
		}
		util := dsl.NewDSLUtil()
		result, err := util.Json.Map.Replace(text, jsonMapReplacement)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(result)
	},
}
