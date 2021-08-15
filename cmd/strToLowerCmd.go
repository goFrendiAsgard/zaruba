package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var strToLowerCmd = &cobra.Command{
	Use:   "toLower <string>",
	Short: "Turn string into lower case",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		fmt.Println(strings.ToLower(args[0]))
	},
}
