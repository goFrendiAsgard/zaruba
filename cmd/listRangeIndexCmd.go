package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listRangeIndexCmd = &cobra.Command{
	Use:   "rangeIndex <list>",
	Short: "Print list indexes",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		list := []interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(cmd, logger, decoration, err)
		}
		for i := 0; i < len(list); i++ {
			fmt.Println(i)
		}
	},
}
