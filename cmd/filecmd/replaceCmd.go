package filecmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var replaceLong = `
Replace a file at a particular index with a new content.
The index is started from 0. You can use a negative index to count from the end of the file.
If not specified, the default index will be 0.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍓
🍇
, and you want to replace 🍓 with a 🍕 before .

------------------------------------------------
Elements | Index  | Note
------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- replace this with 🍕
🍇       | 2/-1   |

Then, you need to invoke the following command:
> zaruba file replace \
  fruits.txt \
  🍕 \
  --index=1

The content of "fruits.txt" will be updated into:
🍊
🍕
🍇

`

var replaceExample = `
> cat fruits.txt
🍊
🍓
🍇

> zaruba file replace \
  fruits.txt \
  '🍕'
> cat fruits.txt
🍕
🍓
🍇

> zaruba file replace \
  fruits.txt \
  '🍕' \
  --index=1
> cat fruits.txt
🍊
🍕
🍇
`

var replaceIndex *int
var replaceCmd = &cobra.Command{
	Use:     "replace <strFileName> <strNewContent>",
	Short:   "Replace a file at a particular index with a new content",
	Long:    replaceLong,
	Example: replaceExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, strNewContent := args[0], args[1]
		if err := util.File.ReplaceLineAtIndex(strFileName, strNewContent, *replaceIndex); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
