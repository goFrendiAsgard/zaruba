package filecmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getLineIndexLong = `
Getting line index of a line that match the last element of the pattern.
Index is started from 0. You can use negative index to count from the end of the file.
`

var getLineIndexExample = `
> cat myFile.txt
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
`
var getLineIndexDesiredPatternIndex *int
var getLineIndexCmd = &cobra.Command{
	Use:     "getLineIndex <strFileName> <jsonStrListPatterns>",
	Short:   "Get line index from a file content matching the pattern",
	Long:    getLineIndexLong,
	Example: getLineIndexExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, jsonStrListPatterns := args[0], args[1]
		jsonLines, err := util.File.ReadLines(strFileName)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		matchIndex, _, err := util.Json.List.GetLinesSubmatch(jsonLines, jsonStrListPatterns, *getLineIndexDesiredPatternIndex)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(matchIndex)
	},
}
