package envcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var readPrefix *string
var readCmd = &cobra.Command{
	Use:   "read <strFileName>",
	Short: "Read environment variable declarations from environment file as a jsonMap",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		fileName := args[0]
		util := dsl.NewDSLUtil()
		jsonMapEnv, err := util.File.ReadEnv(fileName)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if *readPrefix != "" {
			jsonMapEnv, err = util.Json.Map.CascadePrefixKeys(jsonMapEnv, *readPrefix)
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		}
		fmt.Println(jsonMapEnv)
	},
}
