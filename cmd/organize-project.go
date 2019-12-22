package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/organizer"
)

func init() {
	rootCmd.AddCommand(organizeProjectCmd)
}

var organizeProjectCmd = &cobra.Command{
	Use:   "organize-project [project-dir [...args]]",
	Short: "Organize a project",
	Long:  `Zaruba will manage project-dependency and perform organize-project script in every directory`,
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
		cwd, _ := os.Getwd()
		log.Printf("[INFO] Invoking organize-project. cwd: %s, project-dir: %s, other arguments: %#v", cwd, projectDir, arguments)
		if err := organizer.Organize(projectDir, organizer.NewOption(), arguments...); err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
