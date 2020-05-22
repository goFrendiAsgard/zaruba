package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
	"github.com/state-alchemists/zaruba/modules/organizer"
)

func init() {
	rootCmd.AddCommand(organizeCmd)
}

var organizeCmd = &cobra.Command{
	Use:   "organize [project-dir [...args]]",
	Short: "Organize project.",
	Long:  "Manage component dependencies in a project.",
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
		// get absolute project dir and project config
		projectDir, err := filepath.Abs(projectDir)
		if err != nil {
			logger.Fatal(err)
		}
		p, err := config.CreateProjectConfig(projectDir)
		if err != nil {
			logger.Fatal(err)
		}
		// invoke action
		if err = organizer.Organize(projectDir, p, arguments...); err != nil {
			logger.Fatal(err)
		}
	},
}
