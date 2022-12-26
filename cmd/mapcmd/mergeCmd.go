package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var mergeExample = `
> zaruba map merge '{"server": "localhost", "port": 3306}' '{"protocol": "mysql"}'
{"server": "localhost", "port": 3306, "protocol": "mysql"}
`

var mergeCmd = &cobra.Command{
	Use:     "merge <jsonMap> <otherJsonMaps...>",
	Short:   "Merge multiple jsonMaps",
	Example: mergeExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		jsonMap := args
		util := dsl.NewDSLUtil()
		jsonMapMerged, err := util.Json.Map.Merge(jsonMap...)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(jsonMapMerged)
	},
}
