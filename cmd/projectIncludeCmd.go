package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var projectIncludeCmd = &cobra.Command{
	Use:   "include <projectFile> <fileName>",
	Short: "Add file to project",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		project, err := getProject(decoration, projectFile)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		if err = project.Init(); err != nil {
			exit(cmd, logger, decoration, err)
		}
		fileName := args[1]
		if err = config.IncludeFileToProject(project, fileName); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
