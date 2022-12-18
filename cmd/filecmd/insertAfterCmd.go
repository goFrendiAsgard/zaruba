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

Line  | Index
-------------
a     | 0/-3
b     | 1/-2
c     | 2/-1
`

var insertAfterExample = `
> cat myFile.txt
a
b
c

> zaruba file insertAfter myFile.txt d

> cat myFile.txt
a
b
c
d

> zaruba file insertAfter myFile.txt e --index=0
a
e
b
c
d
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
