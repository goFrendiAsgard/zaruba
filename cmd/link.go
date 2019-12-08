package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/link"
)

func init() {
	rootCmd.AddCommand(linkCmd)
}

var linkCmd = &cobra.Command{
	Use:   "link <project-dir> <source> <destination>",
	Short: "Register dependencies",
	Long:  `Zaruba will register dependency to project's zaruba.dependency.json`,
	Run: func(cmd *cobra.Command, args []string) {
		// handle invalid parameter
		if len(args) < 3 {
			log.Fatal("[ERROR] `project-dir`, `source`, and `destination` expected, current arguments: ", args)
		}
		// get projectDir
		projectDir := args[0]
		source := args[1]
		destination := args[2]
		// invoke action
		cwd, _ := os.Getwd()
		log.Printf("[INFO] Invoking link. cwd: %s, project-dir: %s, source: %s, destination: %s", cwd, projectDir, source, destination)
		link.Create(projectDir, source, destination)
	},
}
