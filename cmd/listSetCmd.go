package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var listSetCmd = &cobra.Command{
	Use:   "set <list> <index> <value>",
	Short: "Set list[index] to value and return new JSON list",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		listString := args[0]
		util := core.NewCoreUtil()
		index, err := strconv.Atoi(args[1])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		value := args[2]
		newListString, err := util.Json.List.Set(listString, index, value)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(newListString)
	},
}
