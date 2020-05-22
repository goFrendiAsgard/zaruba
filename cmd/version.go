package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Zaruba's version.",
	Long:  "Zaruba's version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Zaruba v0.1.0 -- [prototype]")
	},
}
