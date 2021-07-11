package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/output"
)

var readLinesCmd = &cobra.Command{
	Use:   "readLines <fileName>",
	Short: "Read lines from file",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		fileName := args[0]
		list, err := file.ReadLines(fileName)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		resultB, err := json.Marshal(list)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}

func init() {
	rootCmd.AddCommand(readLinesCmd)
}
