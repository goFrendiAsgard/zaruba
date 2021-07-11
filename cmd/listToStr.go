package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listToStrCmd = &cobra.Command{
	Use:   "listToStr <list> [separator]",
	Short: "Transform JSON list into single string",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		list := []string{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		separator := "\n"
		if len(args) > 1 {
			separator = args[1]
		}
		fmt.Println(strings.Join(list, separator))
	},
}

func init() {
	rootCmd.AddCommand(listToStrCmd)
}
