package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(linkCmd)
}

var linkCmd = &cobra.Command{
	Use:   "link <project-dir> <source> <destination>",
	Short: "Register dependencies",
	Long:  `Zaruba will register dependency to zaruba.dependency.json`,
	Run: func(cmd *cobra.Command, args []string) {
		// handle invalid parameter
		if len(args) < 3 {
			log.Fatal("[ERROR] `project-dir`, `source`, and `destination` expected")
		}
		// get projectDir
		projectDir := args[0]
		source := args[1]
		destination := args[2]
		// invoke action
		log.Printf("[INFO] Invoking link. project-dir: %s, source: %s, destination: %s", projectDir, source, destination)
	},
}
