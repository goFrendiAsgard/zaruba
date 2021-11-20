package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var envWriteCmd = &cobra.Command{
	Use:   "write <fileName> <envMap> [prefix]",
	Short: "Write envMap to file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		fileName, mapString := args[0], args[1]
		var err error
		util := core.NewCoreUtil()
		if len(args) > 1 {
			prefix := args[1]
			mapString, err = util.Json.Map.CascadePrefixKeys(mapString, prefix)
			if err != nil {
				exit(cmd, logger, decoration, err)
			}
		}
		envString, err := util.Json.Map.ToEnvString(mapString)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		if err := util.File.WriteText(fileName, envString, 0755); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
