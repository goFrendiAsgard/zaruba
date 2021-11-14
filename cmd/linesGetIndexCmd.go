package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var linesGetIndexCmd = &cobra.Command{
	Use:   "getIndex <list> <patterns>",
	Short: "Return index of matching the pattern",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		list := []string{}
		err := json.Unmarshal([]byte(args[0]), &list)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		patterns := []string{}
		if err := json.Unmarshal([]byte(args[1]), &patterns); err != nil {
			exit(cmd, logger, decoration, err)
		}
		util := core.NewUtil()
		index, _, err := util.Str.GetLineSubmatch(list, patterns)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(index)
	},
}
