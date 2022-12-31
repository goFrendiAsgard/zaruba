package taskcmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var syncEnvCmd = &cobra.Command{
	Use:   "syncEnv <taskName> [projectFile]",
	Short: "Update task's environment",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		taskName := args[0]
		projectFilePath, err := cmdHelper.GetProjectRelFilePath(args, 1, "index.zaruba.yaml", "index.zaruba.yml")
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		util := dsl.NewDSLUtil()
		if err = util.Project.Task.Env.Sync(taskName, projectFilePath); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
