package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var toPascalCaseCmd = &cobra.Command{
	Use:   "toPascalCase <string>",
	Short: "Turn string into PascalCase",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for toPascalCase"))
		}
		fmt.Println(str.ToPascalCase(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(toPascalCaseCmd)
}
