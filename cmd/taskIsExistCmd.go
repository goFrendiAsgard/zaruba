package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var taskIsExistCmd = &cobra.Command{
	Use:   "isExist <projectFile> <taskName>",
	Short: "Is task exist",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		taskName := args[1]
		util := core.NewCoreUtil()
		exist, err := util.Project.Task.IsExist(projectFile, taskName)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		if exist {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
