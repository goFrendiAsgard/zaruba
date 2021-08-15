package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var mapTransformKeyCmd = &cobra.Command{
	Use:   "transformKey <map> <prefix> [suffix]",
	Short: "Transform map keys",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		dict := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &dict); err != nil {
			exit(commandName, logger, decoration, err)
		}
		prefix, suffix := args[1], ""
		if len(args) > 2 {
			suffix = args[2]
		}
		newDict := map[string]interface{}{}
		for key, val := range dict {
			newKey := fmt.Sprintf("%s%s%s", prefix, key, suffix)
			newDict[newKey] = val
		}
		resultB, err := json.Marshal(newDict)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
