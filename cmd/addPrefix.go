package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var addPrefixCmd = &cobra.Command{
	Use:   "addPrefix <string> <prefix>",
	Short: "Add prefix to string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 2 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for addPrefix"))
		}
		if strings.HasPrefix(args[0], args[1]) {
			fmt.Println(args[0])
			return
		}
		fmt.Println(args[1] + args[0])
	},
}

func init() {
	rootCmd.AddCommand(addPrefixCmd)
}
