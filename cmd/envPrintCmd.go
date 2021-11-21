package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var envPrintCmd = &cobra.Command{
	Use:   "print <envMap> [prefix]",
	Short: "Print environment",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		mapString := args[0]
		var err error
		util := core.NewCoreUtil()
		if len(args) > 1 {
			prefix := args[1]
			mapString, err = util.Json.Map.CascadePrefixKeys(mapString, prefix)
			if err != nil {
				exit(cmd, args, logger, decoration, err)
			}
		}
		envString, err := util.Json.Map.ToEnvString(mapString)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(envString)
	},
}
