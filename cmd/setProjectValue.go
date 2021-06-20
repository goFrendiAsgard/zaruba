package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var setProjectValueCmd = &cobra.Command{
	Use:   "setProjectValue",
	Short: "Set project value",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 3 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument"))
		}
		if err := util.SetProjectValue(args[0], args[1], args[2]); err != nil {
			showErrorAndExit(logger, decoration, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setProjectValueCmd)
}
