package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/component"
)

func init() {
	rootCmd.AddCommand(createComponentCmd)
}

var createComponentCmd = &cobra.Command{
	Use:   "create <template> [project-dir]",
	Short: "Create Component",
	Long:  `Zaruba will manage project-dependency and perform organize-project.sh`,
	Run: func(cmd *cobra.Command, args []string) {
		// handle invalid parameter
		if len(args) < 1 {
			log.Fatal("[ERROR] template is expected, current arguments: ", args)
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
		if err := component.Create(template, projectDir, arguments...); err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
