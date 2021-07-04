package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var addProjectLinkCmd = &cobra.Command{
	Use:   "addProjectLink <valueFile> <source> <destination>",
	Short: "Add project link",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 3 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for addProjectLink"))
		}
		if err := util.AddProjectLink(args[0], args[1], args[2]); err != nil {
			showErrorAndExit(logger, decoration, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addProjectLinkCmd)
}
