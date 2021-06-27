package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var capitalizeCmd = &cobra.Command{
	Use:   "capitalize",
	Short: "make capitalized string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for capitalize"))
		}
		fmt.Println(str.Capitalize(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(capitalizeCmd)
}
