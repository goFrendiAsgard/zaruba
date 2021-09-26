package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/utility"
)

var strToCamelCmd = &cobra.Command{
	Use:   "toCamel <string>",
	Short: "Turn string into camelCase",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		util := utility.NewUtil()
		fmt.Println(util.Str.ToCamel(args[0]))
	},
}
