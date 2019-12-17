package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/action"
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
			log.Fatal("[ERROR] action expected, current arguments: ", args)
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
		// invoke action
		cwd, _ := os.Getwd()
		log.Printf("[INFO] Invoking %s. cwd: %s, project-dir: %s, other arguments: %s", cwd, actionString, projectDir, arguments)
		if err := action.Do(actionString, projectDir, arguments...); err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
