package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var strFullIndentCmd = &cobra.Command{
	Use:   "fullIndent <string> <indentation>",
	Short: "indent multi-line string, include first line",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		text := args[0]
		indentation := args[1]
		util := core.NewCoreUtil()
		fmt.Println(util.Str.FullIndent(text, indentation))
	},
}
