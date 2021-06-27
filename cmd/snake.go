package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var snakeCmd = &cobra.Command{
	Use:   "snake",
	Short: "make snake-cased string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for snake"))
		}
		fmt.Println(str.Snake(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(snakeCmd)
}
