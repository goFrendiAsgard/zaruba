package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var listLengthCmd = &cobra.Command{
	Use:   "length <list>",
	Short: "Get list's length",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		list := []interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(len(list))
	},
}
