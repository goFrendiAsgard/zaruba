package projectcmd

import (
	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var syncEnvCmd = &cobra.Command{
	Use:   "syncEnv [projectFile]",
	Short: "Update every task's environment",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 0)
		projectFilePath, err := cmdHelper.GetProjectRelFilePath(args, 0, "index.zaruba.yaml", "index.zaruba.yml")
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		util := dsl.NewDSLUtil()
		if err = util.Project.SyncTasksEnv(projectFilePath); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
