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
		d := logger.NewDecoration()
		logger.Printf("%s%sDon't worry 👌%s, everyone makes mistakes\n", d.Bold, d.Yellow, d.Normal)
	},
}

func init() {
	rootCmd.AddCommand(sorryCmd)
}
