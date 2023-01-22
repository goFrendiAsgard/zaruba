package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var replaceLong = `
Replace a jsonStringList at a particular index with a new lines.
The Index is started from 0. You can use a negative index to count from the end of the jsonStringList.
If not specified, the default index will be 0.

For example, you have a jsonStringList ["🍊", "🍓", "🍇"]
, and you want to replace 🍓 with two 🍕.

-------------------------------------------------
Elements | Index  | Note
-------------------------------------------------
🍊       | 0/-3   |
🍓       | 1/-2   | <-- replace this with two🍕
🍇       | 2/-1   |

In that case, you need to invoke the following command:
> zaruba lines replace \
  '["🍊", "🍓", "🍇"]' \
  '["🍕", "🍕"]' \
  --index=1

The result will be:
["🍊","🍕","🍕","🍇"]
`

var replaceExample = `
> LINES='["🍊", "🍓", "🍇"]'
> zaruba lines replace "$LINES" '🍕'
["🍕","🍓","🍇"]

> LINES='["🍊", "🍓", "🍇"]'
> zaruba lines replace "$LINES" '["🍕", "🍕"]' --index=1
["🍊","🍕","🍕","🍇"]

> LINES='["🍊", "🍓", "🍇"]'
> zaruba lines replace "$LINES" '["🍕"]' --index=-1
["🍊","🍓","🍕"]
`

var replaceIndex *int
var replaceCmd = &cobra.Command{
	Use:     "replace <jsonStrList> <jsonStrListNewLines | strNewLine>",
	Short:   "Replace a jsonStringList at a particular index with new lines",
	Long:    replaceLong,
	Example: replaceExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		jsonStrList, jsonStrListNewContent := args[0], args[1]
		newJsonStrList, err := util.Json.List.ReplaceLineAtIndex(jsonStrList, *replaceIndex, jsonStrListNewContent)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(newJsonStrList)
	},
}
