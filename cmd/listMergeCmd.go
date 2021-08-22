package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listMergeCmd = &cobra.Command{
	Use:   "merge <list> <otherList...>",
	Short: "Merge JSON lists",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		newList := []interface{}{}
		for _, arg := range args {
			list := []interface{}{}
			if err := json.Unmarshal([]byte(arg), &list); err != nil {
				exit(cmd, logger, decoration, err)
			}
			newList = append(newList, list...)
		}
		resultB, err := json.Marshal(newList)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
