package strcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var addPrefixCmd = &cobra.Command{
	Use:   "addPrefix <string> <prefix>",
	Short: "Add prefix to string or do nothing if string already has that prefix",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := core.NewCoreUtil()
		fmt.Println(util.Str.AddPrefix(args[0], args[1]))
	},
}
