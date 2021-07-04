package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var toLowerCaseCmd = &cobra.Command{
	Use:   "toLowerCase <string>",
	Short: "Turn string into lower case",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for toLowerCase"))
		}
		fmt.Println(strings.ToLower(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(toLowerCaseCmd)
}
