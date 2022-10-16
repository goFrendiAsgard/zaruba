package advertisementcmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/state-alchemists/zaruba/advertisement"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var showCmd = &cobra.Command{
	Use:   "show <advertisementFile>",
	Short: "Show advertisement",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		advertisementFile := args[0]
		advs, err := advertisement.NewAdvs(advertisementFile)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(advs.Get())
	},
}
