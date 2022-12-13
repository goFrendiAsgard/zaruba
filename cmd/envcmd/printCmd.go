package envcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var printPrefix *string
var printCmd = &cobra.Command{
	Use:     "print <envMap> [fileName]",
	Short:   "Print environment as json map",
	Aliases: []string{"write"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		mapString := args[0]
		fileName := ""
		if len(args) > 1 {
			fileName = args[1]
		}
		var err error
		util := dsl.NewDSLUtil()
		if *printPrefix != "" {
			mapString, err = util.Json.Map.CascadePrefixKeys(mapString, *printPrefix)
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		}
		envString, err := util.Json.Map.ToEnvString(mapString)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if fileName != "" {
			if err := util.File.WriteText(fileName, envString, 0755); err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		}
		fmt.Println(envString)
	},
}
