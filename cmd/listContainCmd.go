package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listContainCmd = &cobra.Command{
	Use:   "contain <list> <element>",
	Short: "Find out whether list contains string or not",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		list := []interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		var seekElement interface{}
		if err := json.Unmarshal([]byte(args[1]), &seekElement); err != nil {
			seekElement = args[1]
		}
		for _, element := range list {
			if element == seekElement {
				fmt.Println(1)
				return
			}
		}
		fmt.Println(0)
	},
}
