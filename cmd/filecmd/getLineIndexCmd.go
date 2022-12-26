package filecmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getLineIndexLong = `
Return the index of a line matching a particular index at a specified patterns.
Index is started from 0. You can use a negative index to count from the end of the file.
If not specified, the default index will be -1.

For example, you have a file named "fruits.txt" containing the following text:
🍊
🍓
🍇
🍊
🍓
🍇
You want to get the index of an 🍓 that is located after two 🍊 and before a 🍇.

---------------------------------------------------------------------------------
Elements | Element index  | Patterns | Pattern Index | Note
---------------------------------------------------------------------------------
🍊       | 0              | 🍊       | 0/-4          |
🍓       | 1              |          |               |
🍇       | 2              |          |               |
🍊       | 3              | 🍊       | 1/-3          |
🍓       | 4              | 🍓       | 2/-2          | <-- We want this 🍓
🍇       | 5              | 🍇       | 3/-1          |


Then, you need to invoke the following command:
> zaruba file getLineIndex \
  fruits.txt \
  '["🍊", "🍊", "🍓","🍇"]' \
  --index=2

The result will be: 4
`

var getLineIndexExample = `
> cat fruits.txt
🍊
🍓
🍇
🍊
🍓
🍇

> zaruba file getLineIndex \
  fruits.txt \
  '🍓'
1

> zaruba file getLineIndex \
  fruits.txt \
  '["🍊", "🍊", "🍓","🍇"]' \
  --index=1
3

> zaruba file getLineIndex \
  fruits.txt \
  '["🍊", "🍊", "🍓","🍇"]' \
  --index=-1
5
`
var getLineIndexDesiredPatternIndex *int
var getLineIndexCmd = &cobra.Command{
	Use:     "getLineIndex <strFileName> <jsonStrListPatterns>",
	Short:   "Return the index of a line matching a particular index at a specified patterns",
	Long:    getLineIndexLong,
	Example: getLineIndexExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, jsonStrListPatterns := args[0], args[1]
		jsonLines, err := util.File.ReadLines(strFileName)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		matchIndex, _, err := util.Json.List.GetLinesSubmatch(jsonLines, jsonStrListPatterns, *getLineIndexDesiredPatternIndex)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(matchIndex)
	},
}
