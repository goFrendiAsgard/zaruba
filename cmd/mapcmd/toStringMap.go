package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var toStringMapExample = `
> zaruba map toStringMap '{"server": "localhost", "port": 3306, "env": {"enable_ui": 0}}'
{"env":"{\"enable_ui\":0}","port":"3306","server":"localhost"}
`

var toStringMapCmd = &cobra.Command{
	Use:     "toStringMap <jsonMap>",
	Short:   "Transform a jsonMap into a jsonStringMap",
	Example: toStringMapExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonMap := args[0]
		util := dsl.NewDSLUtil()
		newJsonMap, err := util.Json.Map.ToStringMap(jsonMap)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(newJsonMap)
	},
}
