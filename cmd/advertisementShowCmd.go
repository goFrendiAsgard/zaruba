package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/advertisement"
	"github.com/state-alchemists/zaruba/output"
)

var advertisementShowCmd = &cobra.Command{
	Use:   "show <advertisementFile>",
	Short: "Show advertisement",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		advertisementFile := args[0]
		advs, err := advertisement.NewAdvs(advertisementFile)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(advs.Get())
	},
}
