package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "zaruba",
	Short: "Zaruba is technology agnostic artefact generator",
	Long:  `Zaruba will watch over your project. Detect any changes in your files, and perform necessary actions to maintain consistency.`,
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
