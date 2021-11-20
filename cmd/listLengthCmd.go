package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var listLengthCmd = &cobra.Command{
	Use:   "length <list>",
	Short: "Get list's length",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		util := core.NewCoreUtil()
		length, err := util.Json.List.GetLength(args[0])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(length)
	},
}
