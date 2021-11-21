package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var listContainCmd = &cobra.Command{
	Use:   "contain <jsonList> <element>",
	Short: "Find out whether jsonList contains string or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		listString := args[0]
		elementStr := args[1]
		util := core.NewCoreUtil()
		exist, err := util.Json.List.Contain(listString, elementStr)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		if exist {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
