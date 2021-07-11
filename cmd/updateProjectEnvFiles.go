package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var updateProjectEnvFilesCmd = &cobra.Command{
	Use:   "updateProjectEnvFiles <projectFile>",
	Short: "Update every environment files (*.env) in project file's directory based on defined tasks",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			exit(commandName, logger, decoration, fmt.Errorf("too few arguments"))
		}
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		projectDir := filepath.Dir(projectFile)
		csvRecordLogger := getCsvRecordLogger(projectDir)
		project, err := getProject(logger, decoration, csvRecordLogger, projectFile)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		if err = project.Init(); err != nil {
			exit(commandName, logger, decoration, err)
		}
		config.UpdateProjectEnvFiles(project)
	},
}

func init() {
	rootCmd.AddCommand(updateProjectEnvFilesCmd)
}
