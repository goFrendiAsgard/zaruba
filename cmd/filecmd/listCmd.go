package filecmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var listCmd = &cobra.Command{
	Use:   "list <path>",
	Short: "list files/folders in a path",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		dirPath := args[0]
		util := dsl.NewDSLUtil()
		fileNames, err := util.File.List(dirPath)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		for _, fileName := range fileNames {
			fmt.Println(fileName)
		}
	},
}
