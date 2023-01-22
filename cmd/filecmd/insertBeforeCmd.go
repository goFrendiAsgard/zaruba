package filecmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertBeforeLong = `
Insert a new content into a text file before a particular index.
The index is started from 0. You can use a negative index to count from the end of the file.
If not specified, the default index will be 0.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍓
🍇
, and you want to insert a 🍕 before 🍓.

------------------------------------------------
Elements | Index  | Note
------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- insert a 🍕 before this
🍇       | 2/-1   |

In that case, you need to invoke the following command:
> zaruba file insertBefore fruits.txt \
  🍕 \
  --index=1

The content of "fruits.txt" will be updated into:
🍊
🍕
🍓
🍇
`

var insertBeforeExample = `
> echo 🍊 > fruits.txt
> echo 🍓 >> fruits.txt
> echo 🍇 >> fruits.txt
> zaruba file insertBefore fruits.txt 🍕
> cat fruits.txt
🍕
🍊
🍓
🍇
> echo 🍊 > fruits.txt
> echo 🍓 >> fruits.txt
> echo 🍇 >> fruits.txt
> zaruba file insertBefore fruits.txt 🍕 --index=1
> cat fruits.txt
🍊
🍕
🍓
🍇
`

var insertBeforeIndex *int
var insertBeforeCmd = &cobra.Command{
	Use:     "insertBefore <strFileName> <strNewContent>",
	Short:   "Insert a new content into a text file before a particular index",
	Long:    insertBeforeLong,
	Example: insertBeforeExample,
	Aliases: []string{"prepend"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, strNewContent := args[0], args[1]
		if err := util.File.InsertLineBeforeIndex(strFileName, strNewContent, *insertBeforeIndex); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
