package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var strGetIndentationCmd = &cobra.Command{
	Use:   "getIndentation <string> [level=1]",
	Short: "Get indentation of string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		level, err := strconv.Atoi(args[1])
		if err != nil {
			level = 1
		}
		result, err := str.GetSingleIndentation(text, level)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(result)
	},
}
