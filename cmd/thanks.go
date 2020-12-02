package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/logger"
)

// thanksCmd represents the thanks command
var thanksCmd = &cobra.Command{
	Use:     "thanks",
	Short:   "Say thank you to Zaruba",
	Long:    "💀 Say thank you to Zaruba",
	Aliases: []string{"thankYou", "thankyou"},
	Run: func(cmd *cobra.Command, args []string) {
		d := logger.NewDecoration()
		logger.Printf("%s%sYour welcome 😊%s\n", d.Bold, d.Yellow, d.Normal)
		logger.Printf("Please consider donating ☕ to my creator here:\n")
		logger.Printf("%shttps://paypal.me/gofrendi%s\n", d.Yellow, d.Normal)
	},
}

func init() {
	rootCmd.AddCommand(thanksCmd)
}
