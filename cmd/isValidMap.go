package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var isValidMapCmd = &cobra.Command{
	Use:   "isValidMap <value>",
	Short: "Check whether value is valid JSON map or not",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		value := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &value); err != nil {
			fmt.Println(0)
			return
		}
		fmt.Println(1)

	},
}

func init() {
	rootCmd.AddCommand(isValidMapCmd)
}
