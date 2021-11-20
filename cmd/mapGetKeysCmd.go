package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var mapGetKeysCmd = &cobra.Command{
	Use:   "getKeys <jsonMap>",
	Short: "Return JSON string list containing keys of JSON map",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		mapString := args[0]
		util := core.NewCoreUtil()
		keys, err := util.Json.Map.GetKeys(mapString)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(util.Json.FromInterface(keys))
	},
}
