package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var generateCmd = &cobra.Command{
	Use:   "generate <templateLocation> <destination> <replacementMap>",
	Short: "Make something based on template",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		sourceTemplatePath := args[0]
		destinationPath := args[1]
		mapString := args[2]
		util := core.NewCoreUtil()
		replacementMap, err := util.Json.Map.GetStringDict(mapString)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		if err := util.File.Generate(sourceTemplatePath, destinationPath, replacementMap); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
