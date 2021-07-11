package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listToMapCmd = &cobra.Command{
	Use:   "listToMap <list> <keys>",
	Short: "Transform JSON list into JSON map",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		list := []string{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		keys := []string{}
		if err := json.Unmarshal([]byte(args[1]), &keys); err != nil {
			exit(commandName, logger, decoration, err)
		}
		result := map[string]string{}
		for index, key := range keys {
			value := list[index]
			result[key] = value
		}
		resultB, err := json.Marshal(result)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}

func init() {
	rootCmd.AddCommand(listToMapCmd)
}
