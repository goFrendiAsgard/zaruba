package filecmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getLineLong = `
Getting a line that match the last element of the pattern.
Index is started from 0. You can use negative index to count from the end of the file.

Line                          | Index
-------------------------------------------
class Num:                    | 0/-5
    def __init__(self, num):  | 1/-4
        self.num = num        | 2/-3
    def add(self, addition):  | 3/-2
        self.num += addition  | 4/-1
`

var getLineExample = `
> cat num.py
class Num:
    def __init__(self, num):
        self.num = num
    def add(self, addition):
        self.num += addition

> zaruba file getLine num.py 0
class Num:

> zaruba file getLine num.py 2
        self.num = num

> zaruba file getLine num.py -1
        self.num += addition
`
var getLineCmd = &cobra.Command{
	Use:     "getLine <strFileName> <index>",
	Short:   "Return desired line of a file content",
	Long:    getLineLong,
	Example: getLineExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, indexStr := args[0], args[1]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		jsonLines, err := util.File.ReadLines(strFileName)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		stringList, err := util.Json.ToStringList(jsonLines)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if index < 0 {
			index = len(stringList) + index
		}
		fmt.Println(stringList[index])
	},
}
