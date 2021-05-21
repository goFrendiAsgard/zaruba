package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

// thanksCmd represents the thanks command
var thanksCmd = &cobra.Command{
	Use:     "thanks",
	Short:   "Say thank you to Zaruba",
	Long:    "💀 Say thank you to Zaruba",
	Aliases: []string{"thankYou", "thankyou"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		output.ShowThanksResponse(logger, decoration)
	},
}

func init() {
	rootCmd.AddCommand(thanksCmd)
}
