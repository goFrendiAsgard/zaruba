package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var showAdvCmd = &cobra.Command{
	Use:   "showAdv",
	Short: "Show advertisement",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for showAdv"))
		}
		message, err := util.GetAdv(args[0])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		fmt.Println(message)
	},
}

func init() {
	rootCmd.AddCommand(showAdvCmd)
}
