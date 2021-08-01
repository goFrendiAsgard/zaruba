package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var syncTaskEnvCmd = &cobra.Command{
	Use:   "syncTaskEnv <projectFile> <taskName>",
	Short: "Update task's environment",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
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
		task, taskExist := project.Tasks[args[1]]
		if !taskExist {
			exit(commandName, logger, decoration, fmt.Errorf("task %s does not exist", args[1]))
		}
		if err = config.SyncTaskEnv(task); err != nil {
			exit(commandName, logger, decoration, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(syncTaskEnvCmd)
}
