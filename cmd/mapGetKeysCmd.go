package cmd

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var mapGetKeysCmd = &cobra.Command{
	Use:   "getKeys <map>",
	Short: "Return JSON string list containing keys of JSON map",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		dict := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &dict); err != nil {
			exit(cmd, logger, decoration, err)
		}
		result := []string{}
		for key := range dict {
			result = append(result, key)
		}
		sort.Strings(result)
		resultB, err := json.Marshal(result)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
