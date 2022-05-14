package taskcmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	common "github.com/state-alchemists/zaruba/jsonutil/helper"
	"github.com/state-alchemists/zaruba/output"
)

var setEnvsCmd = &cobra.Command{
	Use:     "setConfigs <taskName> {<envMap> | <envKey> <envValue>} [projectFile]",
	Short:   "Set task env",
	Aliases: []string{"setEnv"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := core.NewCoreUtil()
		taskName, jsonEnvMap := args[0], args[1]
		projectFileArgIndex := 2
		envMap, err := util.Json.ToStringDict(jsonEnvMap)
		if err != nil {
			if len(args) <= 2 {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			envMap = common.StringDict{}
			envMap[args[1]] = args[2]
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
		if err = util.Project.Task.Env.Set(taskName, envMap, projectFile); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
