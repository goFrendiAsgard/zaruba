package strcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var toPluralCmd = &cobra.Command{
	Use:   "toPlural <string>",
	Short: "Turn string into PascalCase",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		util := core.NewCoreUtil()
		fmt.Println(util.Str.ToPlural(args[0]))
	},
}
