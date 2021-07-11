package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var addProjectLinkCmd = &cobra.Command{
	Use:   "addProjectLink <valueFile> <source> <destination>",
	Short: "Add project link",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 3)
		if err := config.AddProjectLink(args[0], args[1], args[2]); err != nil {
			exit(commandName, logger, decoration, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addProjectLinkCmd)
}
