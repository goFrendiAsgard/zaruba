package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var getLineSubmatchCmd = &cobra.Command{
	Use:   "getLineSubmatch <list> <patterns>",
	Short: "Return submatch matching the pattern",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		list := []string{}
		err := json.Unmarshal([]byte(args[0]), &list)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		patterns := []string{}
		if err := json.Unmarshal([]byte(args[1]), &patterns); err != nil {
			exit(commandName, logger, decoration, err)
		}
		index, submatch, err := str.GetLineSubmatch(list, patterns)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		if index == -1 {
			exit(commandName, logger, decoration, fmt.Errorf("no line match %#v", patterns))
		}
		resultB, err := json.Marshal(submatch)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}

func init() {
	rootCmd.AddCommand(getLineSubmatchCmd)
}
