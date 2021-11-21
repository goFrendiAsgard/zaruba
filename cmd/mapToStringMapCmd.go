package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var mapToStringMapCmd = &cobra.Command{
	Use:   "toStringMap <jsonMap>",
	Short: "Transform to string map",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		mapString := args[0]
		util := core.NewCoreUtil()
		newMapString, err := util.Json.Map.ToStringMap(mapString)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newMapString)
	},
}
