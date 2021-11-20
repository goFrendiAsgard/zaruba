package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var listGetCmd = &cobra.Command{
	Use:   "get <list> <index>",
	Short: "Get list[index]",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		listString := args[0]
		index, err := strconv.Atoi(args[1])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		util := core.NewCoreUtil()
		data, err := util.Json.List.GetValue(listString, index)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(util.Json.FromInterface(data))
	},
}
