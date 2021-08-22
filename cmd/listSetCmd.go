package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listSetCmd = &cobra.Command{
	Use:   "set <list> <index> <value>",
	Short: "Set list[index] to value and return new JSON list",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		list := []interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(cmd, logger, decoration, err)
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		if index < -1 || index >= len(list) {
			exit(cmd, logger, decoration, fmt.Errorf("index out of bound"))
		}
		var newValue interface{}
		if err := json.Unmarshal([]byte(args[2]), &newValue); err != nil {
			newValue = args[2]
		}
		list[index] = newValue
		resultB, err := json.Marshal(list)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
