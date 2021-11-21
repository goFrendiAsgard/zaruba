package cmd

import (
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
		jsonString, err := util.File.ReadLines(fileName)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonString)
	},
}
