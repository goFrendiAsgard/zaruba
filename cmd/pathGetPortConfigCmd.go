package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var pathGetPortConfigCmd = &cobra.Command{
	Use:   "getPortConfig <location>",
	Short: "Return jsonList representing default configs.ports",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		util := core.NewCoreUtil()
		location := args[0]
		jsonList, err := util.Path.GetPortConfigByLocation(location)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonList)
	},
}
