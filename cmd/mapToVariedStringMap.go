package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var mapToVariedStringMapCmd = &cobra.Command{
	Use:   "toVariedStringMap <jsonMap> [keys...]",
	Short: "Transform to string map",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		mapString, keys := args[0], args[1:]
		util := core.NewCoreUtil()
		newMapString, err := util.Json.Map.ToVariedStringMap(mapString, keys...)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(newMapString)
	},
}
