package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getCmd = &cobra.Command{
	Use:   "get <jsonMap> <strKey>",
	Short: "Get value from JSON map",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		jsonMap, strKey := args[0], args[1]
		util := dsl.NewDSLUtil()
		data, err := util.Json.Map.GetValue(jsonMap, strKey)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(util.Json.FromInterface(data))
	},
}
