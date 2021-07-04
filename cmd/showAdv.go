package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/advertisement"
	"github.com/state-alchemists/zaruba/output"
)

var showAdvCmd = &cobra.Command{
	Use:   "showAdv <advertisementFile>",
	Short: "Show advertisement",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for showAdv"))
		}
		advertisementFile := args[0]
		advs, err := advertisement.NewAdvs(advertisementFile)
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		fmt.Println(advs.Get())
	},
}

func init() {
	rootCmd.AddCommand(showAdvCmd)
}
