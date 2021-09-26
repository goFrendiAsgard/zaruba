package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/utility"
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
		util := utility.NewUtil()
		fmt.Println(util.Str.Indent(text, indentation))
	},
}
