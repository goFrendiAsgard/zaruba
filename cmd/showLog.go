package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var showLogCmd = &cobra.Command{
	Use:   "showLog",
	Short: "Show log",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 2 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument"))
		}
		message, err := util.GetLog(decoration, args[0], args[1])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		fmt.Println(message)
	},
}

func init() {
	rootCmd.AddCommand(showLogCmd)
}
