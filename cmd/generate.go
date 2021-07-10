package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/output"
)

var generateCmd = &cobra.Command{
	Use:   "generate <templateLocation> <destination> <replacementMap>",
	Short: "Make something based on template",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 3 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for generate"))
		}
		templateLocation, err := filepath.Abs(args[0])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		destination, err := filepath.Abs(args[1])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		replacementMap := map[string]string{}
		if err := json.Unmarshal([]byte(args[2]), &replacementMap); err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		file.Generate(templateLocation, destination, replacementMap)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
