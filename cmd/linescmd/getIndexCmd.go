package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getIndexLong = `
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

var getIndexExample = `
> CONTENT='[
"class Num:",
"    def __init__(self, num):",
"        self.num = num",
"    def add(self, addition):",
"        self.num += addition"
]'


> PATTERN='[
"class Num:",
"    def add(self, addition):",
"        self.num += addition"
]'

> zaruba lines getIndex $CONTENT $PATTERN
4

> zaruba lines getIndex $CONTENT $PATTERN --index=-1
4

> zaruba list get $PATTERN 0
class Num:
> zaruba lines getIndex $CONTENT $PATTERN --index=0
0

> zaruba list get $PATTERN 1
    def add(self, addition):
> zaruba lines getIndex $CONTENT $PATTERN --index=1
3
`

var getIndexDesiredPatternIndex *int
var getIndexCmd = &cobra.Command{
	Use:     "getIndex <jsonStrList> <jsonStrListPatterns>",
	Short:   "Return index of lines matching the patterns at desired pattern index",
	Long:    getIndexLong,
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
