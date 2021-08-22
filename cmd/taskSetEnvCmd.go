package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var taskSetEnvCmd = &cobra.Command{
	Use:   "setEnv <projectFile> <taskName> <envMap>",
	Short: "Set task env",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 3)
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		project, err := getProject(decoration, projectFile)
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
		envMap := map[string]string{}
		if err := json.Unmarshal([]byte(args[2]), &envMap); err != nil {
			exit(commandName, logger, decoration, err)
		}
		if err = config.SetTaskEnv(task, envMap); err != nil {
			exit(commandName, logger, decoration, err)
		}
	},
}
