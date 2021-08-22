package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var strToUpperCmd = &cobra.Command{
	Use:   "toUpper <string>",
	Short: "Turn string into UPPER CASE",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		fmt.Println(strings.ToUpper(args[0]))
	},
}
