package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var strEscapeShellArg = &cobra.Command{
	Use:   "escapeShellArg <string>",
	Short: "Escape shell arg",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		util := core.NewCoreUtil()
		fmt.Println(util.Str.EscapeShellArg(text))
	},
}
