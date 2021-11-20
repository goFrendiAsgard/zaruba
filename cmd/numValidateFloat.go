package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var numValidateFloatCmd = &cobra.Command{
	Use:   "validateFloat <value>",
	Short: "Check whether value is valid float or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		if _, err := strconv.ParseFloat(args[0], 64); err != nil {
			fmt.Println(0)
			return
		}
		fmt.Println(1)
	},
}
