package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listSetCmd = &cobra.Command{
	Use:   "listSet <list> <index> <value>",
	Short: "Set list[index] to value and return new JSON string list",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 3)
		list := []string{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		if index < -1 || index >= len(list) {
			exit(commandName, logger, decoration, fmt.Errorf("index out of bound"))
		}
		value := args[2]
		list[index] = value
		resultB, err := json.Marshal(list)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}

func init() {
	rootCmd.AddCommand(listSetCmd)
}
