package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var strSplitCmd = &cobra.Command{
	Use:   "split <string> [separator]",
	Short: "Split string into JSON list",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		separator := "\n"
		if len(args) > 1 {
			separator = args[1]
		}
		util := core.NewCoreUtil()
		result := util.Str.Split(text, separator)
		resultB, err := json.Marshal(result)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))

	},
}
