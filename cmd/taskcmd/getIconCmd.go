package taskcmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var getIconCmd = &cobra.Command{
	Use:   "getIcon <taskName> [projectFile]",
	Short: "get task icon",
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
		util := core.NewCoreUtil()
		icon, err := util.Project.Task.GetIcon(taskName, projectFile)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(icon)
	},
}
