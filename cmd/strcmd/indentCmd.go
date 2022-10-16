package strcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var indentCmd = &cobra.Command{
	Use:   "indent <string> <indentation>",
	Short: "indent multi-line string, exclude first line",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		text := args[0]
		indentation := args[1]
		util := dsl.NewDSLUtil()
		fmt.Println(util.Str.Indent(text, indentation))
	},
}
