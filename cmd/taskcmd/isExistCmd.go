package taskcmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var isExistCmd = &cobra.Command{
	Use:   "isExist <taskName> [projectFile]",
	Short: "Is task exist",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		taskName := args[0]
		projectFile := "index.zaruba.yaml"
		if len(args) > 1 {
			projectFile = args[1]
		}
		projectFile, err := filepath.Abs(projectFile)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		util := dsl.NewDSLUtil()
		exist, err := util.Project.Task.IsExist(taskName, projectFile)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if exist {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	},
}
