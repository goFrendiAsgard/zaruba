package taskcmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	common "github.com/state-alchemists/zaruba/jsonutil/helper"
	"github.com/state-alchemists/zaruba/output"
)

var setConfigsCmd = &cobra.Command{
	Use:     "setConfigs <taskName> {<jsonMapConfig> | <configKey> <configValue>} [projectFile]",
	Short:   "Set task configs",
	Aliases: []string{"setConfig"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		taskName, jsonConfigMap := args[0], args[1]
		projectFileArgIndex := 2
		configMap, err := util.Json.ToStringDict(jsonConfigMap)
		if err != nil {
			if len(args) <= 2 {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
			configMap = common.StringDict{}
			configMap[args[1]] = args[2]
			projectFileArgIndex = 3
		}
		projectFilePath, err := cmdHelper.GetProjectRelFilePath(args, projectFileArgIndex, "index.zaruba.yaml", "index.zaruba.yml")
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		if err = util.Project.Task.Config.Set(taskName, configMap, projectFilePath); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
