package projectcmd

import (
	"path/filepath"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var syncEnvFilesCmd = &cobra.Command{
	Use:   "syncEnvFiles [projectFile]",
	Short: "Update environment files (*.env) in project file's directory",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 0)
		projectFile := "index.zaruba.yaml"
		if len(args) > 0 {
			projectFile = args[0]
		}
		projectFile, err := filepath.Abs(projectFile)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		util := dsl.NewDSLUtil()
		if err := util.Project.SyncEnvFiles(projectFile); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
