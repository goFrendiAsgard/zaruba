package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var addTaskDependencyCmd = &cobra.Command{
	Use:   "addTaskDependency <projectFile> <taskName> <dependencyTaskNames>",
	Short: "Add task dependency",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 3)
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
		taskName := args[1]
		task, taskExist := project.Tasks[taskName]
		if !taskExist {
			exit(commandName, logger, decoration, fmt.Errorf("task %s is not exist", taskName))
		}
		dependencyTaskNames := []string{}
		if err = json.Unmarshal([]byte(args[2]), &dependencyTaskNames); err != nil {
			dependencyTaskNames = []string{args[2]}
		}
		if err = config.AddTaskDependencies(task, dependencyTaskNames); err != nil {
			exit(commandName, logger, decoration, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addTaskDependencyCmd)
}
