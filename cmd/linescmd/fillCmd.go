package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var fillCmd = &cobra.Command{
	Use:   "fill <jsonList> <patterns> <suplements>",
	Short: "Insert suplements to lines if patterns is not found",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		util := dsl.NewDSLUtil()
		jsonLines, jsonPatterns, jsonSuplements := args[0], args[1], args[2]
		jsonNewLines, err := util.Json.List.CompleteLines(jsonLines, jsonPatterns, jsonSuplements)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonNewLines)
	},
}
