package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var pathGetAppNameCmd = &cobra.Command{
	Use:     "getAppName <string>",
	Short:   "Get default app name based on location or image name",
	Aliases: []string{"getServiceName"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		util := core.NewCoreUtil()
		defaultAppName, err := util.Path.GetDefaultAppName(args[0])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(defaultAppName)
	},
}
