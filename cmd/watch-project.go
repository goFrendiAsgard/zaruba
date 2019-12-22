package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/watcher"
)

func init() {
	rootCmd.AddCommand(watchProjectCmd)
}

var watchProjectCmd = &cobra.Command{
	Use:   "watch-project [project-dir [...args]]",
	Short: "Watch and organize a project",
	Long:  `Zaruba will perform "organize" whenever something changed`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir := "."
		arguments := []string{}
		if len(args) > 0 {
			projectDir = args[0]
			if len(args) > 1 {
				arguments = args[1:]
			}
		}
		// invoke action
		errChan := make(chan error)
		go watcher.Watch(projectDir, errChan, make(chan bool), arguments...)
		err := <-errChan
		if err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
