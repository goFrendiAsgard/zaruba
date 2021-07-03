package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var updateProjectTaskEnvCmd = &cobra.Command{
	Use:   "updateProjectTaskEnv <projectFile>",
	Short: "Update any environment files (*.env) in project file's directory by a service name",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for updateProjectTaskEnv"))
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
		if err = util.UpdateProjectTaskEnv(project); err != nil {
			showErrorAndExit(logger, decoration, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateProjectTaskEnvCmd)
}
