package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var replaceLineCmd = &cobra.Command{
	Use:   "replaceLine <lines> <index> <replacement>",
	Short: "Replace lines[index] with replacement",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 3)
		lines := []string{}
		if err := json.Unmarshal([]byte(args[0]), &lines); err != nil {
			exit(commandName, logger, decoration, err)
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		replacements := []string{}
		if err := json.Unmarshal([]byte(args[2]), &replacements); err != nil {
			replacements = []string{args[2]}
		}
		result, err := str.ReplaceLine(lines, index, replacements)
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
	rootCmd.AddCommand(replaceLineCmd)
}
