package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var banner string = `
      _____
     /     \        My name is Zaruba.
    | () () |       I came to be when Garo came to be.
     \  ^  /        And I am forever with him.
      |||||
`
var rootCmd = &cobra.Command{
	Use:   "zaruba <action> [...args]",
	Short: "Developer's Partner",
	Long: strings.Join([]string{
		"Zaruba is agnostic scaffolding tool, task runner, as well as monorepo management.",
		"Zaruba will help you build project faster.",
	}, "\n"),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(banner)
			cmd.Help()
		}
	},
}

// Execute basic action
func Execute() {
	rootCmd.Execute()
}
