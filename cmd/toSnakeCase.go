package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var toSnakeCaseCmd = &cobra.Command{
	Use:   "toSnakeCase <string>",
	Short: "Turn string into snake_case",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for toSnakeCase"))
		}
		fmt.Println(str.ToSnakeCase(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(toSnakeCaseCmd)
}
