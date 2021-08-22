package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var taskIsExistCmd = &cobra.Command{
	Use:   "isExist <projectFile> <taskName>",
	Short: "Is task exist",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
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
		_, taskExist := project.Tasks[taskName]
		if !taskExist {
			fmt.Println(0)
			return
		}
		fmt.Println(1)
	},
}
