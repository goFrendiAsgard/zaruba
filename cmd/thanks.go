package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/monitor"
	"github.com/state-alchemists/zaruba/response"
)

// thanksCmd represents the thanks command
var thanksCmd = &cobra.Command{
	Use:     "thanks",
	Short:   "Say thank you to Zaruba",
	Long:    "💀 Say thank you to Zaruba",
	Aliases: []string{"thankYou", "thankyou"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := monitor.NewDecoration()
		logger := monitor.NewConsoleLogger(decoration)
		response.ShowThanksResponse(logger, decoration)
	},
}

func init() {
	rootCmd.AddCommand(thanksCmd)
}
