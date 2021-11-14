package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var strIndentCmd = &cobra.Command{
	Use:   "indent <string> <indentation>",
	Short: "indent multi-line string, exclude first line",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		text := args[0]
		indentation := args[1]
		util := core.NewUtil()
		fmt.Println(util.Str.Indent(text, indentation))
	},
}
