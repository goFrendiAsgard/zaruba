package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var updateProjectEnvFilesCmd = &cobra.Command{
	Use:   "updateProjectEnvFiles",
	Short: "Update project env files based on service's name",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 2 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for getProjectServiceNames"))
		}
		projectFile, err := filepath.Abs(args[0])
		serviceName := args[1]
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
		util.UpdateProjectEnvFiles(project, serviceName, projectDir)
	},
}

func init() {
	rootCmd.AddCommand(updateProjectEnvFilesCmd)
}
