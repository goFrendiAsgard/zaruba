package filecmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertBeforeLong = `
Insert a new content at a file name before a particular index.
Index is started from 0. You can use negative index to count from the end of the file.
If not specified, default index is 0.

Line  | Index
-------------
a     | 0/-3
b     | 1/-2
c     | 2/-1
`

var insertBeforeExample = `
> cat myFile.txt
a
b
c

> zaruba file insertBefore myFile.txt d

> cat myFile.txt
d
a
b
c

> zaruba file insertBefore myFile.txt e --index=0
e
d
a
b
c
`

var insertBeforeIndex *int
var insertBeforeCmd = &cobra.Command{
	Use:     "insertBefore <strFileName> <strNewContent>",
	Short:   "Insert a new content at a file name before a particular index",
	Long:    insertBeforeLong,
	Example: insertBeforeExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, strNewContent := args[0], args[1]
		if err := util.File.InsertLineBeforeIndex(strFileName, strNewContent, *insertBeforeIndex); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
