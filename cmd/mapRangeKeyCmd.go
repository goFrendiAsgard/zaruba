package cmd

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var mapRangeKeyCmd = &cobra.Command{
	Use:   "rangeKey <map>",
	Short: "Print map keys",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		dict := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &dict); err != nil {
			exit(cmd, logger, decoration, err)
		}
		keys := []string{}
		for key := range dict {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}
