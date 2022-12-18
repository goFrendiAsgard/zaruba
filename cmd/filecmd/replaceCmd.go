package filecmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var replaceLong = `
Replace a new content at a file name at a particular index.
Index is started from 0. You can use negative index to count from the end of the file.
If not specified, default index is 0.

Line                          | Index
-------------------------------------------
class Num:                    | 0/-5
    def __init__(self, num):  | 1/-4
        self.num = num        | 2/-3
    def add(self, addition):  | 3/-2
        self.num += addition  | 4/-1
`

var replaceExample = `
> cat num.py
class Num:
    def __init__(self, num):
        self.num = num
    def add(self, addition):
        self.num += addition

> zaruba file replace num.py 'class Number:"
> cat num.py
class Number:
    def __init__(self, num):
        self.num = num
    def add(self, addition):
        self.num += addition

> zaruba file replace num.py "    def __init__(self, num: int):" --index=1
> zaruba file replace num.py "    def add(self, addition: int):" --index=3
> cat num.py
class Num:
    def __init__(self, num: int):
        self.num = num
    def add(self, addition: int):
        self.num += addition
`

var replaceIndex *int
var replaceCmd = &cobra.Command{
	Use:     "replace <strFileName> <strNewContent>",
	Short:   "Replace a new content at a file name at a particular index",
	Long:    replaceLong,
	Example: replaceExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, strNewContent := args[0], args[1]
		if err := util.File.ReplaceLineAtIndex(strFileName, strNewContent, *replaceIndex); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
