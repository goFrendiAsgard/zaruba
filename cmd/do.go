package cmd

import (
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do <action> [...args]",
	Short: "Do custom action",
	Long:  `Zaruba will perform custom action`,
	Run: func(cmd *cobra.Command, args []string) {
		// handle invalid parameter
		if len(args) < 1 {
			log.Fatal("[ERROR] action")
		}
		// get `action`, `arguments` and `projectDir`
		action := args[0]
		arguments := args[1:]
		projectDir := "."
		// make projectDir absolute
		projectDir, err := filepath.Abs(projectDir)
		if err != nil {
			log.Fatal("[ERROR] ", err)
		}
		// invoke action
		log.Printf("[INFO] Invoking %s. project-dir: %s, other arguments: %#v", action, projectDir, arguments)
	},
}
