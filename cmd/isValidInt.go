package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var isValidIntCmd = &cobra.Command{
	Use:   "isValidInt <value>",
	Short: "Check whether value is valid int or not",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println(0)
			return
		}
		fmt.Println(1)
	},
}

func init() {
	rootCmd.AddCommand(isValidIntCmd)
}
