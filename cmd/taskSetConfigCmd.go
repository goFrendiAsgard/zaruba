package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var taskSetConfigCmd = &cobra.Command{
	Use:   "setConfig <projectFile> <taskName> <configMap>",
	Short: "Set task config",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 3)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		taskName, jsonConfigMap := args[1], args[2]
		util := core.NewCoreUtil()
		if err = util.Project.Task.Config.Set(projectFile, taskName, jsonConfigMap); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
