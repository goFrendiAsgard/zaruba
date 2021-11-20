package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var mapValidateCmd = &cobra.Command{
	Use:   "validate <jsonMap>",
	Short: "Check whether jsonMap is valid JSON map or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		mapString := args[0]
		util := core.NewCoreUtil()
		if util.Json.Map.Validate(mapString) {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
