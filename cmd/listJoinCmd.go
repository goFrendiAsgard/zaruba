package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var listJoinCmd = &cobra.Command{
	Use:   "join <list> [separator]",
	Short: "Transform JSON list into single string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		listString := args[0]
		separator := "\n"
		if len(args) > 1 {
			separator = args[1]
		}
		util := core.NewCoreUtil()
		str, err := util.Json.List.Join(listString, separator)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(str)
	},
}
