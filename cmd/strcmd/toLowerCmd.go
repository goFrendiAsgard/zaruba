package strcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var toLowerCmd = &cobra.Command{
	Use:   "toLower <string>",
	Short: "Turn string into lower case",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		util := core.NewCoreUtil()
		fmt.Println(util.Str.ToLower(args[0]))
	},
}
