package envcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getPrefix *string
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get envmap from currently loaded environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 0)
		util := dsl.NewDSLUtil()
		envMapStr, err := util.Json.Map.GetFromEnv()
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if len(args) == 0 {
			fmt.Println(envMapStr)
			return
		}
		cascadedEnvMapStr, err := util.Json.Map.CascadePrefixKeys(envMapStr, *getPrefix)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(cascadedEnvMapStr)
	},
}
