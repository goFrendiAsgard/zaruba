package strcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var splitCmd = &cobra.Command{
	Use:   "split <string> [separator]",
	Short: "Split string into JSON list",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		separator := "\n"
		if len(args) > 1 {
			separator = args[1]
		}
		util := dsl.NewDSLUtil()
		list := util.Str.Split(text, separator)
		jsonList, err := util.Json.FromStringList(list)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonList)

	},
}
