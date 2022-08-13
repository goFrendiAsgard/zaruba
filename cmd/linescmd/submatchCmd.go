package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var submatchCmd = &cobra.Command{
	Use:   "submatch <jsonList> <patterns>",
	Short: "Return submatch matching the pattern",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonLines, jsonPatterns := args[0], args[1]
		matchIndex, jsonSubmatch, err := util.Json.List.GetLinesSubmatch(jsonLines, jsonPatterns)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if matchIndex == -1 {
			cmdHelper.Exit(cmd, args, logger, decoration, fmt.Errorf("no line match %s", jsonPatterns))
		}
		fmt.Println(jsonSubmatch)
	},
}
