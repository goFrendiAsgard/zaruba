package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getIndexExample = `
Getting line index that match the last element of the pattern
    > zaruba lines getIndex '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]'
    4
    > zaruba lines getIndex '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]' --index=-1
    4

lines:        ["a", "a", "b", "c", "d", "e"]
                0    1    2    3    4    5
                                    ^
                                    line index that match the last index of the pattern
patterns:     ["a",    , "b",      "d"]
                0         1         2
                                    ^
                                    last index of the pattern

Getting line index that match the desired index of the pattern
    > zaruba lines getIndex '["a", "a", "b", "c", "d", "e"]' '["a", "b", "d"]' --index=1
    2

lines:        ["a", "a", "b", "c", "d", "e"]
                0    1    2    3    4    5
                          ^
                          line index that match the desired index of the pattern
patterns:     ["a",    , "b",      "d"]
                0         1         2
                          ^
                          desired index of the pattern
`
var getIndexDesiredPatternIndex *int
var getIndexCmd = &cobra.Command{
	Use:     "getIndex <jsonStrList> <jsonStrListPatterns>",
	Short:   "Return index of lines matching the patterns at desiredPatternIndex",
	Example: getIndexExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListPatterns := args[0], args[1]
		matchIndex, _, err := util.Json.List.GetLinesSubmatch(jsonStrList, jsonStrListPatterns, *getIndexDesiredPatternIndex)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(matchIndex)
	},
}
