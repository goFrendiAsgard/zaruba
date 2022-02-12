package filecmd

import (
	"fmt"

	"github.com/spf13/cobra"

	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var walkCmd = &cobra.Command{
	Use:   "walk <path>",
	Short: "list files/folder in a path, recursively",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		dirPath := args[0]
		util := core.NewCoreUtil()
		fileNames, err := util.File.Walk(dirPath)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		for _, fileName := range fileNames {
			fmt.Println(fileName)
		}
	},
}
