package listcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var containCmd = &cobra.Command{
	Use:   "contain <jsonList> <strElement>",
	Short: "Find out whether jsonList contains string or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		jsonList, strElement := args[0], args[1]
		util := dsl.NewDSLUtil()
		exist, err := util.Json.List.Contain(jsonList, strElement)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if exist {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
