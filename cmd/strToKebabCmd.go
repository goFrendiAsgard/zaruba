package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var strToKebabCmd = &cobra.Command{
	Use:   "toKebab <string>",
	Short: "Turn string into kebab-case",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		fmt.Println(str.ToKebabCase(args[0]))
	},
}
