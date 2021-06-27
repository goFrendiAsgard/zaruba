package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var toUpperCmd = &cobra.Command{
	Use:   "toUpper",
	Short: "make upper-cased string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for toUpper"))
		}
		fmt.Println(strings.ToUpper(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(toUpperCmd)
}
