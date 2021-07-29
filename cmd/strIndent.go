package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var strIndentCmd = &cobra.Command{
	Use:   "strIndent <string> <indentation>",
	Short: "Add prefix to string",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		text := args[0]
		indentation := args[1]
		lines := strings.Split(text, "\n")
		for index, line := range lines {
			if strings.Trim(line, " ") != "" {
				lines[index] = indentation + line
			}
		}
		indentedText := strings.Join(lines, "\n")
		fmt.Println(indentation + indentedText)
	},
}

func init() {
	rootCmd.AddCommand(strIndentCmd)
}
