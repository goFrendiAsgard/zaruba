package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var mapMergeCmd = &cobra.Command{
	Use:   "merge <map> <otherMaps...>",
	Short: "Merge JSON maps",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		newDict := map[string]interface{}{}
		for _, arg := range args {
			dict := map[string]interface{}{}
			if err := json.Unmarshal([]byte(arg), &dict); err != nil {
				exit(commandName, logger, decoration, err)
			}
			for key, val := range dict {
				if _, exist := newDict[key]; !exist {
					newDict[key] = val
				}
			}
		}
		resultB, err := json.Marshal(newDict)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
