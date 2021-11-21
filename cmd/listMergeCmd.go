package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var listMergeCmd = &cobra.Command{
	Use:   "merge <list> <otherList...>",
	Short: "Merge JSON lists",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		listStrings := args
		util := core.NewCoreUtil()
		newListString, err := util.Json.List.Merge(listStrings...)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newListString)
	},
}
