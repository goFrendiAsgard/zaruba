package listcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var appendExample = `
> zaruba list append '[]' 🍊 🍓 🍇
["🍊","🍓","🍇"] 
`

var appendCmd = &cobra.Command{
	Use:     "append <jsonStrList> <strNewValues...>",
	Short:   "Append new values to a jsonList",
	Example: appendExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		jsonStrList := args[0]
		strNewValues := args[1:]
		util := dsl.NewDSLUtil()
		newValues, err := util.Json.List.Append(jsonStrList, strNewValues...)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(newValues)
	},
}
