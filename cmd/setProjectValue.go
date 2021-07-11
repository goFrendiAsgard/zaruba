package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var setProjectValueCmd = &cobra.Command{
	Use:   "setProjectValue <valueFile> <key> <value>",
	Short: "Set project value",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 3)
		if err := config.SetProjectValue(args[0], args[1], args[2]); err != nil {
			exit(commandName, logger, decoration, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setProjectValueCmd)
}
