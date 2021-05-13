package cmd

import (
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/response"
)

// sorryCmd represents the sorry command
var sorryCmd = &cobra.Command{
	Use:   "sorry",
	Short: "Apologize to Zaruba",
	Long:  "💀 Apologize to Zaruba",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		response.ShowSorryResponse(logger, decoration)
	},
}

func init() {
	rootCmd.AddCommand(sorryCmd)
}
