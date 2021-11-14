package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var generateCmd = &cobra.Command{
	Use:   "generate <templateLocation> <destination> <replacementMap>",
	Short: "Make something based on template",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		sourceTemplatePath := args[0]
		destinationPath := args[1]
		rawReplacementMap := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[2]), &rawReplacementMap); err != nil {
			exit(cmd, logger, decoration, err)
		}
		replacementMap := convertToMapString(rawReplacementMap)
		util := core.NewUtil()
		if err := util.File.Generate(sourceTemplatePath, destinationPath, replacementMap); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
