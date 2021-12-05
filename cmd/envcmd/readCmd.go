package envcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var readCmd = &cobra.Command{
	Use:   "read <fileName> [prefix]",
	Short: "Read envmap from file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		fileName := args[0]
		util := core.NewCoreUtil()
		mapString, err := util.File.ReadEnv(fileName)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if len(args) > 1 {
			prefix := args[1]
			mapString, err = util.Json.Map.CascadePrefixKeys(mapString, prefix)
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		}
		fmt.Println(mapString)
	},
}
