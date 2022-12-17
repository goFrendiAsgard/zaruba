package listcmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var setCmd = &cobra.Command{
	Use:   "set <jsonList> <index> <value>",
	Short: "Set list[index] to value and return new JSON list",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		jsonList := args[0]
		util := dsl.NewDSLUtil()
		index, err := strconv.Atoi(args[1])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		value := args[2]
		newListString, err := util.Json.List.Set(jsonList, index, value)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newListString)
	},
}
