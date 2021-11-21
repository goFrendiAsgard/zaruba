package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var strSubmatchCmd = &cobra.Command{
	Use:   "submatch <string> <pattern>",
	Short: "Return submatch of string based on pattern",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		text, pattern := args[0], args[1]
		util := core.NewCoreUtil()
		jsonSubmatch, err := util.Json.List.Submatch(text, pattern)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonSubmatch)
	},
}
