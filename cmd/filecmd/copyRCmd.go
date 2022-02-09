package filecmd

import (
	"github.com/spf13/cobra"

	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var copyCmd = &cobra.Command{
	Use:   "copy <source> <destination>",
	Short: "copy files/folders recursively",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		source, destination := args[0], args[1]
		util := core.NewCoreUtil()
		if err := util.File.CopyR(source, destination); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
