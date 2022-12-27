package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var toVariedStringMapExample = `
zaruba map toVariedStringMap '{"server": "localhost", "port": 3306}'
{"\"port\"":"\"3306\"","\"server\"":"\"localhost\"","'port'":"'3306'","'server'":"'localhost'","PORT":"3306","Port":"3306","SERVER":"LOCALHOST","Server":"Localhost","port":"3306","server":"localhost"}
`

var toVariedStringMapCmd = &cobra.Command{
	Use:     "toVariedStringMap <jsonMap> [keys...]",
	Short:   "Transform a jsonMap into a jsonStringMap, every keys and values are transformed into multiple variations",
	Example: toVariedStringMapExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonMap, keys := args[0], args[1:]
		util := dsl.NewDSLUtil()
		newJsonMap, err := util.Json.Map.ToVariedStringMap(jsonMap, keys...)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(newJsonMap)
	},
}
