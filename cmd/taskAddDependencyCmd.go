package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var taskAddDependencyCmd = &cobra.Command{
	Use:   "addDependency <projectFile> <taskName> <dependencyTaskNames>",
	Short: "Add task dependency",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
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
		taskName := args[1]
		task, taskExist := project.Tasks[taskName]
		if !taskExist {
			exit(cmd, logger, decoration, fmt.Errorf("task %s is not exist", taskName))
		}
		dependencyTaskNames := []string{}
		if err = json.Unmarshal([]byte(args[2]), &dependencyTaskNames); err != nil {
			dependencyTaskNames = []string{args[2]}
		}
		if err = core.AddTaskDependencies(task, dependencyTaskNames); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
