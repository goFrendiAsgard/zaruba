package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var setCmd = &cobra.Command{
	Use:   "set <jsonMap> <key> <value> [<otherKey> <otherValue>...]",
	Short: "Set map[key] to value",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		mapString := args[0]
		setArgs := args[1:]
		util := core.NewCoreUtil()
		newMapString, err := util.Json.Map.Set(mapString, setArgs...)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newMapString)
	},
}
