package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getKeysCmd = &cobra.Command{
	Use:   "getKeys <jsonMap>",
	Short: "Return JSON string list containing keys of JSON map",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		mapString := args[0]
		util := dsl.NewDSLUtil()
		keys, err := util.Json.Map.GetKeys(mapString)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(util.Json.FromStringList(keys))
	},
}
