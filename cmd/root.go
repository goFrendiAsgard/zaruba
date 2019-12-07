package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "zaruba",
	Short: "Zaruba is agnostic generator and task runner",
	Long:  `Zaruba will help you create project and maintain dependencies among components`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("You call zaruba?")
		log.Println("Try `zaruba help`")
	},
}

// Execute basic action
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
