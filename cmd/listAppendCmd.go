package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var listAppendCmd = &cobra.Command{
	Use:   "append <list> <newValues...>",
	Short: "Append new values to list",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		listString, value := args[0], args[1]
		util := core.NewCoreUtil()
		newListString, err := util.Json.List.Append(listString, value)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(newListString)
	},
}
