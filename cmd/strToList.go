package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var strToListCmd = &cobra.Command{
	Use:   "strToList <string> [separator]",
	Short: "Split string into JSON list",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		text := args[0]
		separator := "\n"
		if len(args) > 1 {
			separator = args[1]
		}
		result := strings.Split(text, separator)
		resultB, err := json.Marshal(result)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))

	},
}

func init() {
	rootCmd.AddCommand(strToListCmd)
}
