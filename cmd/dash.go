package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var dashCmd = &cobra.Command{
	Use:   "dash <string>",
	Short: "Make dashed string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for dash"))
		}
		fmt.Println(str.Dash(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(dashCmd)
}
