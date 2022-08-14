package listcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var validateCmd = &cobra.Command{
	Use:   "validate <jsonList>",
	Short: "Check whether jsonList is valid JSON list or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		listString := args[0]
		util := dsl.NewDSLUtil()
		if util.Json.List.Validate(listString) {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
