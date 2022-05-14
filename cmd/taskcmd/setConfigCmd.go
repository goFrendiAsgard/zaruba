package taskcmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	common "github.com/state-alchemists/zaruba/jsonutil/helper"
	"github.com/state-alchemists/zaruba/output"
)

var setConfigsCmd = &cobra.Command{
	Use:     "setConfigs <taskName> {<configMap> | <configKey> <configValue>} [projectFile]",
	Short:   "Set task configs",
	Aliases: []string{"setConfig"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := core.NewCoreUtil()
		taskName, jsonConfigMap := args[0], args[1]
		projectFileArgIndex := 2
		configMap, err := util.Json.ToStringDict(jsonConfigMap)
		if err != nil {
			if len(args) <= 2 {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			configMap = common.StringDict{}
			configMap[args[1]] = args[2]
			projectFileArgIndex = 3
		}
		projectFile := "index.zaruba.yaml"
		if len(args) > projectFileArgIndex {
			projectFile = args[projectFileArgIndex]
		}
		projectFile, err = filepath.Abs(projectFile)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if err = util.Project.Task.Config.Set(taskName, configMap, projectFile); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
