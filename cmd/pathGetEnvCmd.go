package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var pathGetEnvCmd = &cobra.Command{
	Use:   "getEnv <location>",
	Short: "Return JSON string map containing environment variables defined on location",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		util := core.NewCoreUtil()
		jsonMap, err := util.Path.GetEnvByLocation(args[0])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonMap)
	},
}
