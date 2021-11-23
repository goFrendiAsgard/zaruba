package strcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var submatchCmd = &cobra.Command{
	Use:   "submatch <string> <pattern>",
	Short: "Return submatch of string based on pattern",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		text, pattern := args[0], args[1]
		util := core.NewCoreUtil()
		submatch, err := util.Str.Submatch(text, pattern)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		jsonSubmatch, err := util.Json.FromStringList(submatch)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonSubmatch)
	},
}
