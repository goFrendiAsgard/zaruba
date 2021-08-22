package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var numValidateIntCmd = &cobra.Command{
	Use:   "validateInt <value>",
	Short: "Check whether value is valid int or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println(0)
			return
		}
		fmt.Println(1)
	},
}
