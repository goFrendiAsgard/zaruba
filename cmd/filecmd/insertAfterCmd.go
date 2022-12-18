package filecmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertAfterLong = `
Insert a new content at a file name after a particular index.
Index is started from 0. You can use negative index to count from the end of the file.
If not specified, default index is -1.

Line                          | Index
-------------------------------------------
class Num:                    | 0/-5
    def __init__(self, num):  | 1/-4
        self.num = num        | 2/-3
    def add(self, addition):  | 3/-2
        self.num += addition  | 4/-1
`

var insertAfterExample = `
> cat num.py
class Num:
    def __init__(self, num):
        self.num = num
    def add(self, addition):
        self.num += addition

> zaruba file insertAfter num.py '# EOF"
> cat num.py
class Num:
    def __init__(self, num):
        self.num = num
    def add(self, addition):
        self.num += addition
# EOF

> zaruba file insertAfter num.py "    '''A numeric class'''" --index=0
> cat num.py
class Num:
	'''A numeric class'''
    def __init__(self, num):
        self.num = num
    def add(self, addition):
        self.num += addition
`

var insertAfterIndex *int
var insertAfterCmd = &cobra.Command{
	Use:     "insertAfter <strFileName> <strNewContent>",
	Short:   "Insert a new content at a file name after a particular index",
	Long:    insertAfterLong,
	Example: insertAfterExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, strNewContent := args[0], args[1]
		if err := util.File.InsertLineAfterIndex(strFileName, strNewContent, *insertAfterIndex); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
