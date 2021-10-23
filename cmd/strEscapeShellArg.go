package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/utility"
)

var strEscapeShellArg = &cobra.Command{
	Use:   "escapeShellArg <string>",
	Short: "Escape shell arg",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		util := utility.NewUtil()
		fmt.Println(util.Str.EscapeShellArg(text))
	},
}
