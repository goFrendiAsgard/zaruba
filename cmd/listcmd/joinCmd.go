package listcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var joinExample = `
> zaruba list join '["🍊","🍓","🍇"]' ";"
🍊;🍓;🍇
`

var joinCmd = &cobra.Command{
	Use:     "join <jsonList> [separator]",
	Short:   "Transform a jsonList into single string",
	Example: joinExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonList, separator := args[0], "\n"
		if len(args) > 1 {
			separator = args[1]
		}
		util := dsl.NewDSLUtil()
		str, err := util.Json.List.Join(jsonList, separator)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(str)
	},
}
