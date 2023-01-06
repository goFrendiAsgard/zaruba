package jsoncmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getExample = `
> zaruba json get '{"characters": [{"name": "doraemon"}, {"name": "nobita"}]}' 'characters.[0].name'
"doraemon"
`

var getCmd = &cobra.Command{
	Use:     "get <jsonAny> <key>",
	Short:   "Get value of nested JsonMap or JsonList",
	Example: getExample,
	Aliases: []string{"write"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		jsonString, key := args[0], args[1]
		util := dsl.NewDSLUtil()
		val, err := util.Json.Get(jsonString, key)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(val)
	},
}
