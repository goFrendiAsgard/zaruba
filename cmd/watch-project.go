package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/watcher"
)

func init() {
	rootCmd.AddCommand(watchProjectCmd)
}

var watchProjectCmd = &cobra.Command{
	Use:   "watch-project [project-dir] [...args]",
	Short: "Watch and organize a project",
	Long:  `Zaruba will perform "organize" whenever something changed`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir := "."
		if len(args) > 0 {
			projectDir = args[0]
		}
		arguments := args[1:]
		// invoke action
		cwd, _ := os.Getwd()
		log.Printf("[INFO] Invoking watch-project. cwd: %s, project-dir: %s, other arguments: %#v", cwd, projectDir, arguments)
		if err := watcher.Watch(projectDir, make(chan bool), arguments...); err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
