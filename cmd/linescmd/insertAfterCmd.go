package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertAfterLong = `
Insert new lines into a jsonStringList after a particular index.
The index is started from 0. You can use a negative index to count from the end of the jsonStringList.
If not specified, the default index will be -1.

For example, you have a jsonStringList ["🍊", "🍓", "🍇"]
, and you want to insert two 🍕 after 🍓.

------------------------------------------------
Elements | Index  | Note
------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- insert two🍕 after this
🍇       | 2/-1   |

Then, you need to invoke the following command:
> zaruba lines insertAfter \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1

The result will be:
["🍊","🍓","🍕","🍕","🍇"]
`

var insertAfterExample = `
> zaruba lines insertAfter \
  '["🍊", "🍓", "🍇"]' \
  '🍕'
["🍊","🍓","🍇", "🍕"]

> zaruba lines insertAfter \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1
["🍊","🍓","🍕","🍕","🍇"]

> zaruba lines insertAfter \
  '["🍊", "🍓", "🍇"]' \
  '["🍕"]' \
  --index=-1
["🍊","🍓","🍇","🍕"]
`

var insertAfterIndex *int
var insertAfterCmd = &cobra.Command{
	Use:     "insertAfter <jsonStrList> <jsonStrListNewLines | strNewLine>",
	Short:   "Insert a new lines into jsonStringList after a particular index",
	Long:    insertAfterLong,
	Example: insertAfterExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListNewLines := args[0], args[1]
		newJsonStrList, err := util.Json.List.InsertLineAfterIndex(jsonStrList, *insertAfterIndex, jsonStrListNewLines)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newJsonStrList)
	},
}
