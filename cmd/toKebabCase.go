package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var toKebabCaseCmd = &cobra.Command{
	Use:   "toKebabCase <string>",
	Short: "Turn string into kebab-case",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for toKebabCase"))
		}
		fmt.Println(str.ToKebabCase(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(toKebabCaseCmd)
}
