package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/organizer"
	"log"
)

func init() {
	rootCmd.AddCommand(organizeCmd)
}

var organizeCmd = &cobra.Command{
	Use:   "organize [project]",
	Short: "Organize a project",
	Long:  `Zaruba will perform necessary action to organize your project`,
	Run: func(cmd *cobra.Command, args []string) {
		// get project
		project := "."
		if len(args) > 0 {
			project = args[0]
		}
		// watch and throw error if necessary
		if err := organizer.Organize(project); err != nil {
			log.Fatal(err)
		}
	},
}
