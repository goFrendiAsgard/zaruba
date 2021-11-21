package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var envGetCmd = &cobra.Command{
	Use:   "get [prefix]",
	Short: "Get envmap from currently loaded environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 0)
		util := core.NewCoreUtil()
		envMapStr, err := util.Json.Map.GetFromEnv()
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		if len(args) == 0 {
			fmt.Println(envMapStr)
			return
		}
		prefix := args[0]
		cascadedEnvMapStr, err := util.Json.Map.CascadePrefixKeys(envMapStr, prefix)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(cascadedEnvMapStr)
	},
}
