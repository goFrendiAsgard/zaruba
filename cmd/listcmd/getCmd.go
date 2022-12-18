package listcmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getExample = `
> zaruba list get '["strawberry", "orange", "apple"]' 0
strawberry

> zaruba list contain '["strawberry", "orange", "apple"]' 2
apple
`

var getCmd = &cobra.Command{
	Use:     "get <jsonList> <index>",
	Short:   "Get jsonList element at an index",
	Example: getExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		jsonList := args[0]
		index, err := strconv.Atoi(args[1])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		util := dsl.NewDSLUtil()
		data, err := util.Json.List.GetValue(jsonList, index)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(util.Json.FromInterface(data))
	},
}
