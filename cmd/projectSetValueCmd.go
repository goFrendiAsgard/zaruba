package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var projectSetValueCmd = &cobra.Command{
	Use:   "setValue <valueFile> <key> <value>",
	Short: "Set project value",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		if err := core.SetProjectValue(args[0], args[1], args[2]); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
