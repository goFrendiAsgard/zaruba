package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var toUpperCaseCmd = &cobra.Command{
	Use:   "toUpperCase <string>",
	Short: "Turn string into UPPER CASE",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for toUpperCase"))
		}
		fmt.Println(strings.ToUpper(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(toUpperCaseCmd)
}
