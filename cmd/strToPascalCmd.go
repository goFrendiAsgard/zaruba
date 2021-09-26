package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/utility"
)

var strToPascalCmd = &cobra.Command{
	Use:   "toPascal <string>",
	Short: "Turn string into PascalCase",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		util := utility.NewUtil()
		fmt.Println(util.Str.ToPascal(args[0]))
	},
}
