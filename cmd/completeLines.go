package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var completeLinesCmd = &cobra.Command{
	Use:   "completeLines <lines> <patterns> <suplements>",
	Short: "Insert suplements to lines if patterns is not found",
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
		suplements := []string{}
		if err := json.Unmarshal([]byte(args[2]), &suplements); err != nil {
			suplements = []string{args[2]}
		}
		result, err := str.CompleteLines(lines, patterns, suplements)
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
	rootCmd.AddCommand(completeLinesCmd)
}
