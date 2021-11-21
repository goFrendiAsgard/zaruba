package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var mapGetCmd = &cobra.Command{
	Use:   "get <jsonMap> <key>",
	Short: "Get value from JSON map",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		mapString := args[0]
		key := args[1]
		util := core.NewCoreUtil()
		data, err := util.Json.Map.GetValue(mapString, key)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(util.Json.FromInterface(data))
	},
}
