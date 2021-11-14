package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var strAddPrefixCmd = &cobra.Command{
	Use:   "addPrefix <string> <prefix>",
	Short: "Add prefix to string or do nothing if string already has that prefix",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		util := core.NewCoreUtil()
		fmt.Println(util.Str.AddPrefix(args[0], args[1]))
	},
}
