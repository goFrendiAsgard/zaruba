package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getIndexLong = `
Return the index of a line matching a particular index at a specified patterns.
Index is started from 0. You can use a negative index to count from the end of the jsonStringList.
If not specified, the default index will be -1.

For example, you have a jsonStringList ["🍊", "🍓", "🍇","🍊", "🍓","🍇"].
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


In that case, you need to invoke the following command:
> zaruba lines getIndex \
  '["🍊", "🍓", "🍇","🍊", "🍓","🍇"]' \
  '["🍊", "🍊", "🍓","🍇"]' \
  --index=2

The result will be: 4
`

var getIndexExample = `
> LINES='["🍊", "🍓", "🍇", "🍊", "🍓", "🍇"]'
> zaruba lines getIndex "$LINES" '🍓'
1
> zaruba lines getIndex "$LINES" '["🍊", "🍊", "🍓","🍇"]' --index=1
3
> zaruba lines getIndex "$LINES" '["🍊", "🍊", "🍓","🍇"]' --index=-1
5
`

var getIndexDesiredPatternIndex *int
var getIndexCmd = &cobra.Command{
	Use:     "getIndex <jsonStrList> <jsonStrListPatterns>",
	Short:   "Return the index of a line matching a particular index at a specified patterns",
	Long:    getIndexLong,
	Example: getIndexExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListPatterns := args[0], args[1]
		matchIndex, _, err := util.Json.List.GetLinesSubmatch(jsonStrList, jsonStrListPatterns, *getIndexDesiredPatternIndex)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(matchIndex)
	},
}
