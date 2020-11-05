package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/logger"
)

// sorryCmd represents the sorry command
var sorryCmd = &cobra.Command{
	Use:   "sorry",
	Short: "Apologize to Zaruba",
	Long:  "💀 Apologize to Zaruba",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Printf("Don't worry, everyone makes mistakes\n")
	},
}

func init() {
	rootCmd.AddCommand(sorryCmd)
}
