package filecmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getLineLong = `
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
You want to get the index of a line containing a 🍓 that is located after two 🍊 and before a 🍇.

---------------------------------------------------------------------------------
Elements | Element index  | Patterns | Pattern Index | Note
---------------------------------------------------------------------------------
🍊       | 0              | 🍊       | 0/-4          |
🍓A      | 1              |          |               |
🍇       | 2              |          |               |
🍊       | 3              | 🍊       | 1/-3          |
🍓B      | 4              | 🍓       | 2/-2          | <-- We want this 🍓
🍇       | 5              | 🍇       | 3/-1          |


In that case, you need to invoke the following command:
> zaruba file getLineIndex fruits.txt \
  '["🍊", "🍊", "🍓.*","🍇"]' \
  --index=2

The result will be:
🍓B
`

var getLineExample = `
> echo 🍊A > fruits.txt
> echo 🍓B >> fruits.txt
> echo 🍇C >> fruits.txt
> echo 🍊D >> fruits.txt
> echo 🍓E >> fruits.txt
> echo 🍇F >> fruits.txt
> zaruba file getLineIndex fruits.txt '🍓.*'
🍓B
> zaruba file getLineIndex fruits.txt '["🍊.*", "🍊.*", "🍓.*","🍇.*"]' --index=1
🍊D
> zaruba file getLineIndex fruits.txt '["🍊.*", "🍊.*", "🍓.*","🍇.*"]' --index=-1
🍇F
`

var getLineCmd = &cobra.Command{
	Use:     "getLine <strFileName> <index>",
	Short:   "Return a line matching a particular index at a specified patterns",
	Long:    getLineLong,
	Example: getLineExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, indexStr := args[0], args[1]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		jsonLines, err := util.File.ReadLines(strFileName)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		stringList, err := util.Json.ToStringList(jsonLines)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		if index < 0 {
			index = len(stringList) + index
		}
		fmt.Println(stringList[index])
	},
}
