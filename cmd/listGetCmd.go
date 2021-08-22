package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listGetCmd = &cobra.Command{
	Use:   "get <list> <index>",
	Short: "Get list[index]",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
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
		printInterface(list[index])
	},
}
