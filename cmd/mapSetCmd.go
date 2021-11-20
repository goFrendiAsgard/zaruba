package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var mapSetCmd = &cobra.Command{
	Use:   "set <jsonMap> <key> <value> [<otherKey> <otherValue>...]",
	Short: "Set map[key] to value",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		mapString := args[0]
		setArgs := args[1:]
		util := core.NewCoreUtil()
		newMapString, err := util.Json.Map.Set(mapString, setArgs...)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(newMapString)
	},
}
