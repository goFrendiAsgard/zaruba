package filecmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var submatchLong = `
Return submatch matching the pattern at a desired pattern index.
Index is started from 0. You can use negative index to count from the end of line.

For example, you have a jsonStringList ["🍊", "🍌🍓🍈", "🍇","🍊", "🥑🍓🍎🍏","🍇"].
First, you want to get a line containing a 🍓 that is located after two 🍊 and before a 🍇.
Then you want to get what characters are preceeding/following the 🍓 at that particular line.

---------------------------------------------------------------------------------------------
Elements   | Element index  | Patterns   | Pattern Index | Note
---------------------------------------------------------------------------------------------
🍊         | 0              | 🍊         | 0/-4          |
🍌🍓🍈     | 1              |            |               |
🍇         | 2              |            |               |
🍊         | 3              | 🍊         | 1/-3          |
🥑🍓🍎🍏   | 4              | (.*)🍓(.*) | 2/-2          | <-- We want "🥑" and "🍎🍏"
🍇         | 5              | 🍇         | 3/-1          |

To do this, you need to invoke the following command:
> zaruba lines submatch \
  '["🍊", "🍌🍓🍈", "🍇","🍊", "🥑🍓🍎🍏","🍇"]' \
  '["🍊", "🍊", "(.*)🍓(.*)", "🍇"]' \
  --index=2

The result will be:
["🥑🍓🍎🍏","🥑","🍎🍏"]

You can see that there are three elements of the result:
- The whole line : 🥑🍓🍎🍏
- The characters preceding 🍓: 🥑
- The characters following 🍓: 🍎🍏
`

var submatchExample = `
> zaruba lines submatch \
  '["🍊", "🍌🍓🍈", "🍇","🍊", "🥑🍓🍎🍏","🍇"]' \
  '["🍊", "🍊", "(.*)🍓(.*)", "🍇"]' \
  --index=2
["🥑🍓🍎🍏","🥑","🍎🍏"]
`

var submatchPatternIndex *int
var submatchCmd = &cobra.Command{
	Use:     "submatch <jsonStrList> <jsonStrListPatterns>",
	Short:   "Return submatch matching the pattern at a desired pattern index",
	Long:    submatchLong,
	Example: submatchExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListPatterns := args[0], args[1]
		matchIndex, jsonSubmatch, err := util.File.GetLinesSubmatch(jsonStrList, jsonStrListPatterns, *submatchPatternIndex)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if matchIndex == -1 {
			cmdHelper.Exit(cmd, args, logger, decoration, fmt.Errorf("no line match %s", jsonStrListPatterns))
		}
		fmt.Println(jsonSubmatch)
	},
}
