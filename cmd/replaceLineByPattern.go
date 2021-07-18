package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var replaceLineByPatternCmd = &cobra.Command{
	Use:   "replaceLineByPattern <lines> <patterns> <replacement>",
	Short: "Sequentially match the patterns and replace the first submatch of the last pattern with replacements",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 3)
		lines := []string{}
		if err := json.Unmarshal([]byte(args[0]), &lines); err != nil {
			exit(commandName, logger, decoration, err)
		}
		patterns := []string{}
		if err := json.Unmarshal([]byte(args[1]), &patterns); err != nil {
			exit(commandName, logger, decoration, err)
		}
		replacements := []string{}
		if err := json.Unmarshal([]byte(args[2]), &replacements); err != nil {
			replacements = []string{args[2]}
		}
		result, err := str.ReplaceLineByPattern(lines, patterns, replacements)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		resultB, err := json.Marshal(result)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}

func init() {
	rootCmd.AddCommand(replaceLineByPatternCmd)
}
