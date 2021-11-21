package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var linesSubmatchCmd = &cobra.Command{
	Use:   "submatch <jsonList> <patterns>",
	Short: "Return submatch matching the pattern",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		util := core.NewCoreUtil()
		jsonLines, jsonPatterns := args[0], args[1]
		matchIndex, jsonSubmatch, err := util.Json.List.GetLinesSubmatch(jsonLines, jsonPatterns)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		if matchIndex == -1 {
			exit(cmd, args, logger, decoration, fmt.Errorf("no line match %s", jsonPatterns))
		}
		fmt.Println(util.Json.FromInterface(jsonSubmatch))
	},
}
