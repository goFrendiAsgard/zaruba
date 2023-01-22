package filecmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var replaceAtLong = `
Replace a file at a particular index with a new content.
The index is started from 0. You can use a negative index to count from the end of the file.
If not specified, the default index will be 0.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍓
🍇
, and you want to replace 🍓 with a 🍕.

------------------------------------------------
Elements | Index  | Note
------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- replace this with 🍕
🍇       | 2/-1   |

In that case, you need to invoke the following command:
> zaruba file replaceAt fruits.txt \
  🍕 \
  --index=1

The content of "fruits.txt" will be updated into:
🍊
🍕
🍇
`

var replaceAtExample = `
> echo 🍊 > fruits.txt
> echo 🍓 >> fruits.txt
> echo 🍇 >> fruits.txt
> zaruba file replaceAt fruits.txt 🍕
> cat fruits.txt
🍕
🍓
🍇

> echo 🍊 > fruits.txt
> echo 🍓 >> fruits.txt
> echo 🍇 >> fruits.txt
> zaruba file replaceAt fruits.txt 🍕 --index=1
> cat fruits.txt
🍊
🍕
🍇
`

var replaceAtIndex *int
var replaceAtCmd = &cobra.Command{
	Use:     "replaceAt <strFileName> <strNewContent>",
	Short:   "Replace a file at a particular index with a new content",
	Long:    replaceAtLong,
	Example: replaceAtExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, strNewContent := args[0], args[1]
		if err := util.File.ReplaceLineAtIndex(strFileName, strNewContent, *replaceAtIndex); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
