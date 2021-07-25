package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var indentLinesCmd = &cobra.Command{
	Use:   "indentLines <list> <indentation>",
	Short: "indent lines",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		list := []string{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		indentation := args[1]
		for index, element := range list {
			if strings.Trim(element, " ") != "" {
				list[index] = indentation + element
			}
		}
		resultB, err := json.Marshal(list)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}

func init() {
	rootCmd.AddCommand(indentLinesCmd)
}
