package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/action"
	"github.com/state-alchemists/zaruba/modules/logger"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do <action> [project-dir [...args]]",
	Short: "Do custom action",
	Long:  `Zaruba will perform custom action`,
	Run: func(cmd *cobra.Command, args []string) {
		// handle invalid parameter
		if len(args) < 1 {
			logger.Fatal("action expected, current arguments: ", args)
		}
		// get `action`, `arguments` and `projectDir`
		actionString := args[0]
		projectDir := "."
		arguments := []string{}
		if len(args) >= 2 {
			projectDir = args[1]
			if len(args) > 2 {
				arguments = args[2:]
			}
		}
		arguments = append([]string{projectDir}, arguments...)
		// invoke action
		if err := action.Do(actionString, action.NewOption().SetWorkDir(projectDir), arguments...); err != nil {
			logger.Fatal(err)
		}
	},
}
