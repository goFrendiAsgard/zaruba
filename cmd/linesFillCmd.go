package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var linesFillCmd = &cobra.Command{
	Use:   "fill <jsonList> <patterns> <suplements>",
	Short: "Insert suplements to lines if patterns is not found",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		util := core.NewCoreUtil()
		jsonLines, jsonPatterns, jsonSuplements := args[0], args[1], args[2]
		jsonNewLines, err := util.Json.List.CompleteLines(jsonLines, jsonPatterns, jsonSuplements)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonNewLines)
	},
}
