package taskcmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var setConfigCmd = &cobra.Command{
	Use:   "setConfig <taskName> <configMap> [projectFile]",
	Short: "Set task config",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		taskName, jsonConfigMap := args[0], args[1]
		projectFile := "index.zaruba.yaml"
		if len(args) > 2 {
			projectFile = args[2]
		}
		projectFile, err := filepath.Abs(projectFile)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		util := core.NewCoreUtil()
		if err = util.Project.Task.Config.Set(taskName, jsonConfigMap, projectFile); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
