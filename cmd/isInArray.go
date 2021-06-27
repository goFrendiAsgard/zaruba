package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var isInArrayCmd = &cobra.Command{
	Use:   "isInArray <needle> <haystack> <separator>",
	Short: "Check whether an array contains an element or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 3 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for isInArray"))
		}
		found := util.IsInArray(args[0], args[1], args[2])
		if found {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	},
}

func init() {
	rootCmd.AddCommand(isInArrayCmd)
}
