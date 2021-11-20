package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var linesReadCmd = &cobra.Command{
	Use:   "read <fileName>",
	Short: "Read lines from file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		fileName := args[0]
		util := core.NewCoreUtil()
		list, err := util.File.ReadLines(fileName)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		resultB, err := json.Marshal(list)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
