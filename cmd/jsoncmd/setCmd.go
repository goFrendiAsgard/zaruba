package jsoncmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var setExample = `
> zaruba json set '{"characters": [{"name": "doraemon"}, {"name": "nobita"}]}' 'characters.[1].name' '"dorami"'
{"characters": [{"name": "doraemon"}, {"name": "dorami"}]}
`

var setCmd = &cobra.Command{
	Use:     "set <jsonAny> <key> <value>",
	Short:   "Set value of nested JsonMap or JsonList",
	Example: setExample,
	Aliases: []string{"write"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		jsonString, key, val := args[0], args[1], args[2]
		util := dsl.NewDSLUtil()
		newJsonString, err := util.Json.Set(jsonString, key, val)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(newJsonString)
	},
}
