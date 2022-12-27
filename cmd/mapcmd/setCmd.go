package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var setExample = `
> zaruba list set '{"server": "localhost", "port": 3306}' server stalchmst.com
{"server": "stalchmst.com", "port": 3306}
`

var setCmd = &cobra.Command{
	Use:     "set <jsonMap> <key> <value> [<otherKey> <otherValue>...]",
	Short:   "Set a value of a jsonMap on a particular key",
	Example: setExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		jsonMap, setArgs := args[0], args[1:]
		util := dsl.NewDSLUtil()
		newJsonMap, err := util.Json.Map.Set(jsonMap, setArgs...)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(newJsonMap)
	},
}
