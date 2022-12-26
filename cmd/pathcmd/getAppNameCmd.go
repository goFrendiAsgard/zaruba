package pathcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var getAppNameCmd = &cobra.Command{
	Use:     "getAppName <string>",
	Short:   "Get default app name based on location or image name",
	Aliases: []string{"getServiceName"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		util := dsl.NewDSLUtil()
		defaultAppName, err := util.Path.GetDefaultAppName(args[0])
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(defaultAppName)
	},
}
