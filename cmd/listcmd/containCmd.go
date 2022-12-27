package listcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var containExample = `
> zaruba list contain '["🍊","🍓","🍇"]' 🍓
1

> zaruba list contain '["🍊","🍓","🍇"]' 🍕
0
`

var containCmd = &cobra.Command{
	Use:     "contain <jsonList> <strElement>",
	Short:   "Find out whether a jsonList contains an element or not",
	Example: containExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		jsonList, strElement := args[0], args[1]
		util := dsl.NewDSLUtil()
		exist, err := util.Json.List.Contain(jsonList, strElement)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		if exist {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
