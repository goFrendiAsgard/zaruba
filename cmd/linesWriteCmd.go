package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var linesWriteCmd = &cobra.Command{
	Use:   "write <fileName> <list>",
	Short: "Write list to file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		fileName := args[0]
		util := core.NewCoreUtil()
		list := []string{}
		if err := json.Unmarshal([]byte(args[1]), &list); err != nil {
			exit(cmd, logger, decoration, err)
		}
		if err := util.File.WriteLines(fileName, list, 0755); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
