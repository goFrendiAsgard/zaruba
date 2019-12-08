package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(organizeProjectCmd)
}

var organizeProjectCmd = &cobra.Command{
	Use:   "organize-project [project-dir]",
	Short: "Organize a project",
	Long:  `Zaruba will manage project-dependency and perform organize-project.sh`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir := "."
		if len(args) > 0 {
			projectDir = args[0]
		}
		// invoke action
		log.Printf("[INFO] Invoking organize-project. project-dir: %s", projectDir)
	},
}
