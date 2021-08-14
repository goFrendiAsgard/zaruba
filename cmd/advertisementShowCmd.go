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
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		advertisementFile := args[0]
		advs, err := advertisement.NewAdvs(advertisementFile)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(advs.Get())
	},
}
