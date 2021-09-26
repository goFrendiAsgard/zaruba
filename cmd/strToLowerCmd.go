package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/utility"
)

var strToLowerCmd = &cobra.Command{
	Use:   "toLower <string>",
	Short: "Turn string into lower case",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		util := utility.NewUtil()
		fmt.Println(util.Str.ToLower(args[0]))
	},
}
