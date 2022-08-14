package strcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var toUpperCmd = &cobra.Command{
	Use:   "toUpper <string>",
	Short: "Turn string into UPPER CASE",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		util := dsl.NewDSLUtil()
		fmt.Println(util.Str.ToUpper(args[0]))
	},
}
