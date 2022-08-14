package envcmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var writeCmd = &cobra.Command{
	Use:   "write <fileName> <envMap> [prefix]",
	Short: "Write envMap to file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		fileName, mapString := args[0], args[1]
		var err error
		util := dsl.NewDSLUtil()
		if len(args) > 1 {
			prefix := args[1]
			mapString, err = util.Json.Map.CascadePrefixKeys(mapString, prefix)
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		}
		envString, err := util.Json.Map.ToEnvString(mapString)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if err := util.File.WriteText(fileName, envString, 0755); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
