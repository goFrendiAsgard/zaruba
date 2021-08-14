package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listJoinCmd = &cobra.Command{
	Use:   "join <list> [separator]",
	Short: "Transform JSON list into single string",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		list := []interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		lines := []string{}
		for _, element := range list {
			lines = append(lines, fmt.Sprintf("%v", element))
		}
		separator := "\n"
		if len(args) > 1 {
			separator = args[1]
		}
		fmt.Println(strings.Join(lines, separator))
	},
}
