package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(createComponentCmd)
}

var createComponentCmd = &cobra.Command{
	Use:   "create-component <template> [project-dir]",
	Short: "Create Component",
	Long:  `Zaruba will manage project-dependency and perform organize-project.sh`,
	Run: func(cmd *cobra.Command, args []string) {
		// handle invalid parameter
		if len(args) < 1 {
			log.Fatal("[ERROR] template is expected")
		}
		// get `template`, `projectDir` and `arguments`
		template := args[0]
		projectDir := "."
		arguments := []string{}
		if len(args) > 1 {
			projectDir = args[1]
			arguments = args[2:]
		}
		// invoke action
		log.Printf("[INFO] Invoking organize-project, template: %s. project-dir: %s, other arguments: %#v", template, projectDir, arguments)
	},
}
