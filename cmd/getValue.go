package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var getValueCmd = &cobra.Command{
	Use:   "getValue <value> <defaultValue>",
	Short: "Get value or default value",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		value := args[0]
		if strings.Trim(value, " ") != "" {
			fmt.Println(value)
			return
		}
		defaultValue := args[1]
		fmt.Println(defaultValue)
	},
}

func init() {
	rootCmd.AddCommand(getValueCmd)
}
