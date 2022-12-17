package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var validateCmd = &cobra.Command{
	Use:   "validate <jsonMap>",
	Short: "Check whether jsonMap is valid JSON map or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonMap := args[0]
		util := dsl.NewDSLUtil()
		if util.Json.Map.Validate(jsonMap) {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
