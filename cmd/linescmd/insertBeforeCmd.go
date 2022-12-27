package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var insertBeforeLong = `
Insert new lines into a jsonStringList before a particular index.
The Index is started from 0. You can use a negative index to count from the end of the jsonStringList.
If not specified, the default index will be 0.

For example, you have a jsonStringList ["🍊", "🍓", "🍇"]
, and you want to insert two 🍕 before 🍓.

--------------------------------------------------
Elements | Index  | Note
--------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- insert two🍕 before this
🍇       | 2/-1   |

Then, you need to invoke the following command:
> zaruba lines insertBefore \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1

The result will be:
["🍊","🍕","🍕","🍓","🍇"]
`

var insertBeforeExample = `
> zaruba lines insertBefore \
  '["🍊", "🍓", "🍇"]' \
  '🍕'
["🍕","🍊","🍓","🍇"]

> zaruba lines insertBefore \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1
["🍊","🍕","🍕","🍓","🍇"]

> zaruba lines insertBefore \
  '["🍊", "🍓", "🍇"]' \
  '["🍕"]' \
  --index=-1
["🍊","🍓","🍕","🍇"]
`

var insertBeforeIndex *int
var insertBeforeCmd = &cobra.Command{
	Use:     "insertBefore <jsonStrList> <jsonStrListNewLines | strNewLine>",
	Short:   "Insert new lines into a jsonStringList before a particular index",
	Long:    insertBeforeLong,
	Example: insertBeforeExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListNewLines := args[0], args[1]
		newJsonStrList, err := util.Json.List.InsertLineBeforeIndex(jsonStrList, *insertBeforeIndex, jsonStrListNewLines)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(newJsonStrList)
	},
}
