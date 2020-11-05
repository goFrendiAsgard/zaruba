package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/logger"
)

// thanksCmd represents the thanks command
var thanksCmd = &cobra.Command{
	Use:   "thanks",
	Short: "Say thank you to Zaruba",
	Long:  "💀 Say thank you to Zaruba",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Printf("Your welcome, please consider donating to my creator\n")
	},
}

func init() {
	rootCmd.AddCommand(thanksCmd)
}
