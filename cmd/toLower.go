package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var toLowerCmd = &cobra.Command{
	Use:   "toLower",
	Short: "make lower-cased string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for toLower"))
		}
		fmt.Println(strings.ToLower(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(toLowerCmd)
}
