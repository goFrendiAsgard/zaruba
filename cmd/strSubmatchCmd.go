package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/utility"
)

var strSubmatchCmd = &cobra.Command{
	Use:   "submatch <string> <pattern>",
	Short: "Return submatch of string based on pattern",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		text := args[0]
		pattern := args[1]
		util := utility.NewUtil()
		result, err := util.Str.Submatch(text, pattern)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		resultB, err := json.Marshal(result)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
