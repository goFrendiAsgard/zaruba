package pathcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var getPortConfigCmd = &cobra.Command{
	Use:   "getPortConfig <location>",
	Short: "Return jsonList representing default configs.ports",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		util := dsl.NewDSLUtil()
		location := args[0]
		jsonList, err := util.Path.GetPortConfigByLocation(location)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonList)
	},
}
