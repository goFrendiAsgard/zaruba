package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var strToUpperSnakeCmd = &cobra.Command{
	Use:   "toUpperSnake <string>",
	Short: "Turn string into UPPER_CASE",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		util := core.NewCoreUtil()
		fmt.Println(util.Str.ToUpperSnake(args[0]))
	},
}
