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
	Use:   "organize [project-dir [...args]]",
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
		if err := organizer.Organize(projectDir, organizer.NewOption(), arguments...); err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
