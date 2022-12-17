package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var submatchExample = `
Getting line index that match the last element of the pattern
    > zaruba lines submatch '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]'
    ["d"]
    > zaruba lines submatch '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]' --index=-1
    ["d"]

lines:        ["a", "a", "b", "c", "d", "e"]
                0    1    2    3    4    5
                                    ^
                                    line index that match the last index of the pattern
patterns:     ["a",    , "b",      "d"]
                0         1         2
                                    ^
                                    last index of the pattern

Getting line index that match the desired index of the pattern
    > zaruba lines submatch '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]' --index=1
    ["b"]

lines:        ["a", "a", "b", "c", "d", "e"]
                0    1    2    3    4    5
                          ^
                          line index that match the desired index of the pattern
patterns:     ["a",    , "b",      "d"]
                0         1         2
                          ^
                          desired index of the pattern
`
var submatchDesiredPatternIndex *int
var submatchCmd = &cobra.Command{
	Use:     "submatch <jsonStrList> <jsonStrListPatterns>",
	Short:   "Return submatch matching the pattern",
	Example: submatchExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListPatterns := args[0], args[1]
		matchIndex, jsonSubmatch, err := util.Json.List.GetLinesSubmatch(jsonStrList, jsonStrListPatterns, *submatchDesiredPatternIndex)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if matchIndex == -1 {
			cmdHelper.Exit(cmd, args, logger, decoration, fmt.Errorf("no line match %s", jsonStrListPatterns))
		}
		fmt.Println(jsonSubmatch)
	},
}
