package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var submatchLong = `
Getting line index of a a line that match the last element of the pattern.
Index is started from 0. You can use negative index to count from the end of line

Line                          | Index
-------------------------------------------
class Num:                    | 0/-5
    def __init__(self, num):  | 1/-4
        self.num = num        | 2/-3
    def add(self, addition):  | 3/-2
        self.num += addition  | 4/-1
`

var submatchExample = `
> CONTENT='[
"class Num:",
"    def __init__(self, num):",
"        self.num = num",
"    def add(self, addition):",
"        self.num += addition"
]'

> PATTERN='[
"class Num:",
"( *)def add\\(self, (.*)\\):",
"( *)self\\.num \\+= (.*)"
]'

> zaruba lines submatch $CONTENT $PATTERN
["        self.num += addition","        ","addition"]

> zaruba lines submatch $CONTENT $PATTERN --index=-1
["        self.num += addition","        ","addition"]

> zaruba list get $PATTERN 0
class Num:
> zaruba lines submatch $CONTENT $PATTERN --index=0
["class Num:"]

> zaruba list get $PATTERN 1
( *)def add\(self, (.*)\):
> zaruba lines submatch $CONTENT $PATTERN --index=1
["    def add(self, addition):","    ","addition"]
`

var submatchDesiredPatternIndex *int
var submatchCmd = &cobra.Command{
	Use:     "submatch <jsonStrList> <jsonStrListPatterns>",
	Short:   "Return submatch matching the pattern at desired pattern index",
	Long:    submatchLong,
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
