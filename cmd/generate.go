package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/output"
)

var generateCmd = &cobra.Command{
	Use:   "generate <templateLocation> <destination> <replacementMap>",
	Short: "Make something based on template",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 3)
		sourceTemplatePath := args[0]
		destinationPath := args[1]
		replacementMap := map[string]string{}
		if err := json.Unmarshal([]byte(args[2]), &replacementMap); err != nil {
			exit(commandName, logger, decoration, err)
		}
		file.Generate(sourceTemplatePath, destinationPath, replacementMap)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
