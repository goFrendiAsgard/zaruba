package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listAppendCmd = &cobra.Command{
	Use:   "append <list> <newValues...>",
	Short: "Append new values to list",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		list := []interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		newStrValues := args[1:]
		for _, newStrValue := range newStrValues {
			var newValue interface{}
			if err := json.Unmarshal([]byte(newStrValue), &newValue); err != nil {
				newValue = newStrValue
			}
			list = append(list, newValue)
		}
		resultB, err := json.Marshal(list)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
