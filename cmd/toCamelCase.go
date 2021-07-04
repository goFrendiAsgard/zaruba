package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var toCamelCaseCmd = &cobra.Command{
	Use:   "toCamelCase <string>",
	Short: "Turn string into camelCase",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for toCamelCase"))
		}
		fmt.Println(str.ToCamelCase(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(toCamelCaseCmd)
}
