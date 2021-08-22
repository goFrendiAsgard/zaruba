package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var taskSetConfigCmd = &cobra.Command{
	Use:   "setConfig <projectFile> <taskName> <configMap>",
	Short: "Set task config",
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
		configMap := map[string]string{}
		if err := json.Unmarshal([]byte(args[2]), &configMap); err != nil {
			exit(cmd, logger, decoration, err)
		}
		if err = config.SetTaskConfig(task, configMap); err != nil {
			exit(cmd, logger, decoration, err)
		}
	},
}
