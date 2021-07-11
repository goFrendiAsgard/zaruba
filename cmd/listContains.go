package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listContainsCmd = &cobra.Command{
	Use:   "listContains <list> <string>",
	Short: "Transform JSON list into single string",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		list := []string{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		text := args[1]
		for _, element := range list {
			if strings.Trim(element, " ") == text {
				fmt.Println(1)
				return
			}
		}
		fmt.Println(0)
	},
}

func init() {
	rootCmd.AddCommand(listContainsCmd)
}
