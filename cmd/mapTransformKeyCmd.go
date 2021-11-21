package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var mapTransformKeyCmd = &cobra.Command{
	Use:   "transformKey <jsonMap> <prefix> [suffix]",
	Short: "Transform map keys",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		mapString, prefix, suffix := args[0], args[1], ""
		if len(args) > 2 {
			suffix = args[2]
		}
		util := core.NewCoreUtil()
		newMapString, err := util.Json.Map.TransformKeys(mapString, prefix, suffix)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newMapString)
	},
}
