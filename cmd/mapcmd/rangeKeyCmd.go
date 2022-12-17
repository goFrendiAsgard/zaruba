package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var rangeKeyCmd = &cobra.Command{
	Use:   "rangeKey <jsonMap>",
	Short: "Print map keys",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonMap := args[0]
		util := dsl.NewDSLUtil()
		keys, err := util.Json.Map.GetKeys(jsonMap)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}
