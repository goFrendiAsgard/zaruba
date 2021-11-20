package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var listValidateCmd = &cobra.Command{
	Use:   "validate <value>",
	Short: "Check whether value is valid JSON list or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		listString := args[0]
		util := core.NewCoreUtil()
		if util.Json.List.Validate(listString) {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
