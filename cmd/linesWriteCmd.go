package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var linesWriteCmd = &cobra.Command{
	Use:   "write <fileName> <jsonList>",
	Short: "Write list to file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		fileName, jsonString := args[0], args[1]
		util := core.NewCoreUtil()
		if err := util.File.WriteLines(fileName, jsonString, 0755); err != nil {
			exit(cmd, args, logger, decoration, err)
		}
	},
}
