package filecmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertAfterLong = `
Insert a new content into a text file after a particular index.
The index is started from 0. You can use a negative index to count from the end of the file.
If not specified, the default index will be -1.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍓
🍇
, and you want to insert a 🍕 after 🍓.

------------------------------------------------
Elements | Index  | Note
------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- insert a 🍕 after this
🍇       | 2/-1   |

Then, you need to invoke the following command:
> zaruba file insertAfter \
  fruits.txt \
  🍕 \
  --index=1

The content of "fruits.txt" will be updated into:
🍊
🍓
🍕
🍇
`

var insertAfterExample = `
> cat fruits.txt
🍊
🍓
🍇

> zaruba file insertAfter \
  fruits.txt \
  '🍕'
> cat fruits.txt
🍊
🍓
🍇
🍕

> zaruba file insertAfter \
  fruits.txt \
  '🍕' \
  --index=1
> cat fruits.txt
🍊
🍓
🍕
🍇
`

var insertAfterIndex *int
var insertAfterCmd = &cobra.Command{
	Use:     "insertAfter <strFileName> <strNewContent>",
	Short:   "Insert a new content into a text file after a particular index",
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
