package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var updateProjectEnvFilesCmd = &cobra.Command{
	Use:   "updateProjectEnvFiles <projectFile>",
	Short: "Update every environment files (*.env) in project file's directory based on defined tasks",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for updateProjectEnvFiles"))
		}
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		projectDir := filepath.Dir(projectFile)
		csvRecordLogger := getCsvRecordLogger(projectDir)
		project, err := getProject(logger, decoration, csvRecordLogger, projectFile)
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		if err = project.Init(); err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		util.UpdateProjectEnvFiles(project)
	},
}

func init() {
	rootCmd.AddCommand(updateProjectEnvFilesCmd)
}
