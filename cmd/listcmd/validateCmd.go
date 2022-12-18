package listcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var validateExample = `
> zaruba list validate '["strawberry", "orange", "apple"]'
1

> zaruba list validate 'not a list'
0
`

var validateCmd = &cobra.Command{
	Use:     "validate <jsonList>",
	Short:   "Check whether a jsonList is valid or not",
	Example: validateExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonList := args[0]
		util := dsl.NewDSLUtil()
		if util.Json.List.Validate(jsonList) {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
