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

Line  | Index
-------------
a     | 0/-3
b     | 1/-2
c     | 2/-1
`

var replaceExample = `
> cat myFile.txt
a
b
c

> zaruba file replace myFile.txt d

> cat myFile.txt
d
b
c

> zaruba file replace myFile.txt e --index=2
d
b
e
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
