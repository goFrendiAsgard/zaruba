package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var listRangeIndexCmd = &cobra.Command{
	Use:   "rangeIndex <list>",
	Short: "Print list indexes",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		util := core.NewCoreUtil()
		length, err := util.Json.List.GetLength(args[0])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		for i := 0; i < length; i++ {
			fmt.Println(i)
		}
	},
}
