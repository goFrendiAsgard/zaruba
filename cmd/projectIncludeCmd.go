package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var projectIncludeCmd = &cobra.Command{
	Use:   "include <projectFilePath> <fileName>",
	Short: "Add file to project",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		projectFilePath, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fileName := args[1]
		if err = core.IncludeFileToProject(projectFilePath, fileName); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
