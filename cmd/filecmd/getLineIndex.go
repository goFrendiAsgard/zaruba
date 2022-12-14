package filecmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getLineIndexExample = `
> zaruba file read myFile.txt
a
a
b
c
d
e

> zaruba file getLineIndex myFile.txt '["a", "b", "d"]'
4
> zaruba file getLineIndex myFile.txt '["a", "b", "d"]' --index=-1
4

    Getting line index that match the last element of the pattern
    lines:        ["a", "a", "b", "c", "d", "e"]
                    0    1    2    3    4    5
                                        ^
                                        line index that match the last index of the pattern
    patterns:     ["a",    , "b",      "d"]
                    0         1         2
                                        ^
                                        last index of the pattern


> zaruba lines getLineIndex myFile.txt '["a", "b", "d"]' --index=1
2

    Getting line index that match the desired index of the pattern
    lines:        ["a", "a", "b", "c", "d", "e"]
                    0    1    2    3    4    5
                              ^
                              line index that match the desired index of the pattern
    patterns:     ["a",    , "b",      "d"]
                    0         1         2
                              ^
                              desired index of the pattern
`
var getLineIndexDesiredPatternIndex *int
var getLineIndexCmd = &cobra.Command{
	Use:     "getLineIndex <fileName> <jsonLinesPattern>",
	Short:   "Return index of lines matching the patterns at the file",
	Example: getLineIndexExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		fileName, jsonPatterns := args[0], args[1]
		jsonLines, err := util.File.ReadLines(fileName)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		matchIndex, _, err := util.Json.List.GetLinesSubmatch(jsonLines, jsonPatterns, *getLineIndexDesiredPatternIndex)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(matchIndex)
	},
}
