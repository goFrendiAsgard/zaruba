package projectcmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var includeCmd = &cobra.Command{
	Use:   "include <fileName> [projectFile]",
	Short: "Add file to project",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		projectFilePath := "index.zaruba.yaml"
		fileName := args[0]
		if len(args) > 1 {
			projectFilePath = args[1]
		}
		projectFilePath, err := filepath.Abs(projectFilePath)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		util := core.NewCoreUtil()
		if err = util.Project.IncludeFile(fileName, projectFilePath); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
