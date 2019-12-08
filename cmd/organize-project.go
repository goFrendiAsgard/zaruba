package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/organizer"
)

func init() {
	rootCmd.AddCommand(organizeProjectCmd)
}

var organizeProjectCmd = &cobra.Command{
	Use:   "organize-project [project-dir] [...args]",
	Short: "Organize a project",
	Long:  `Zaruba will manage project-dependency and perform organize-project.sh`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir := "."
		if len(args) > 0 {
			projectDir = args[0]
		}
		arguments := args[1:]
		// invoke action
		log.Printf("[INFO] Invoking organize-project. project-dir: %s, other arguments: %#v", projectDir, arguments)
		if err := organizer.Organize(projectDir, arguments...); err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
