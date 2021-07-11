package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var strToSnakeCmd = &cobra.Command{
	Use:   "strToSnake <string>",
	Short: "Turn string into snake_case",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		fmt.Println(str.ToSnakeCase(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(strToSnakeCmd)
}
