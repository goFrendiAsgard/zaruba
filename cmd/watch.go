package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/watcher"
	"log"
)

func init() {
	rootCmd.AddCommand(watchCmd)
}

var watchCmd = &cobra.Command{
	Use:   "watch [project]",
	Short: "Watch a project",
	Long:  `Zaruba will watch over and organize your project`,
	Run: func(cmd *cobra.Command, args []string) {
		// get project
		project := "."
		if len(args) > 0 {
			project = args[0]
		}
		// watch and throw error if necessary
		if err := watcher.Watch(project, make(chan bool)); err != nil {
			log.Fatal(err)
		}
	},
}
