package filecmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getLineIndexLong = `
Getting line index of a a line that match the last element of the pattern.
Index is started from 0. You can use negative index to count from the end of the file.

Line                          | Index
-------------------------------------------
class Num:                    | 0/-5
    def __init__(self, num):  | 1/-4
        self.num = num        | 2/-3
    def add(self, addition):  | 3/-2
        self.num += addition  | 4/-1
`

var getLineIndexExample = `
> cat num.py
class Num:
    def __init__(self, num):
        self.num = num
    def add(self, addition):
        self.num += addition


> PATTERN='[
"class Num:",
"    def add(self, addition):",
"        self.num += addition"
]'

> zaruba file getLineIndex num.py $PATTERN
4

> zaruba lines getIndex $CONTENT $PATTERN --index=-1
4

> zaruba list get $PATTERN 0
class Num:
> zaruba file getLineIndex $CONTENT $PATTERN --index=0
0

> zaruba list get $PATTERN 1
    def add(self, addition):
> zaruba file getLineIndex $CONTENT $PATTERN --index=1
3
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
