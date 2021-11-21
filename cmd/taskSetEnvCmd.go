package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var taskSetEnvCmd = &cobra.Command{
	Use:   "setEnv <projectFile> <taskName> <envMap>",
	Short: "Set task env",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		taskName, jsonEnvMap := args[1], args[2]
		util := core.NewCoreUtil()
		if err = util.Project.Task.Env.Set(projectFile, taskName, jsonEnvMap); err != nil {
			exit(cmd, args, logger, decoration, err)
		}
	},
}
