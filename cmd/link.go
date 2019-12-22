package cmd

import (
	"log"
	"os"
	"path"
	"path/filepath"

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
		projectDir, _ := filepath.Abs(args[0])
		cwd, _ := os.Getwd()
		source := path.Join(cwd, args[1])
		destination := path.Join(cwd, args[2])
		// invoke action
		log.Printf("[INFO] Invoking link. project-dir: %s, source: %s, destination: %s", projectDir, source, destination)
		link.Create(projectDir, source, destination)
	},
}
