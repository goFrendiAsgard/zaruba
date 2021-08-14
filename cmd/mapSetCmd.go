package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var mapSetCmd = &cobra.Command{
	Use:   "setMapElement <map> <key1> <value1> <key2> <value2>... <keyN> <valueN>",
	Short: "Set map[key] to value",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 3)
		dict := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &dict); err != nil {
			exit(commandName, logger, decoration, err)
		}
		restArgs := args[1:]
		for len(restArgs) > 1 {
			key := restArgs[0]
			var newValue interface{}
			if err := json.Unmarshal([]byte(restArgs[1]), &newValue); err != nil {
				newValue = restArgs[1]
			}
			dict[key] = newValue
			restArgs = restArgs[2:]
		}
		resultB, err := json.Marshal(dict)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
