package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listValidateCmd = &cobra.Command{
	Use:   "validate <value>",
	Short: "Check whether value is valid JSON list or not",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		value := []interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &value); err != nil {
			fmt.Println(0)
			return
		}
		fmt.Println(1)

	},
}
