package listcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var rangeIndexExample = `
> zaruba list rangeIndex '["strawberry", "orange", "apple"]'
0
1
2

> LIST=["strawberry", "orange", "apple"]
> for INDEX in $(zaruba list rangeIndex "$LIST")
  do
	VALUE=$(zaruba list get "$LIST" $INDEX)
	echo "$INDEX $VALUE"
  done

0 strawberry
1 orange
2 apple
`

var rangeIndexCmd = &cobra.Command{
	Use:     "rangeIndex <jsonList>",
	Short:   "Print list indexes",
	Example: rangeIndexExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		util := dsl.NewDSLUtil()
		length, err := util.Json.List.GetLength(args[0])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		for i := 0; i < length; i++ {
			fmt.Println(i)
		}
	},
}
