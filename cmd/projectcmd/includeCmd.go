package projectcmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var includeCmd = &cobra.Command{
	Use:   "include <projectFilePath> <fileName>",
	Short: "Add file to project",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		projectFilePath, err := filepath.Abs(args[0])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fileName := args[1]
		util := core.NewCoreUtil()
		if err = util.Project.IncludeFile(projectFilePath, fileName); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	},
}
