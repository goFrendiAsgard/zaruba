package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var getFromMapCmd = &cobra.Command{
	Use:   "getFromMap <map> <key>",
	Short: "Get value from JSON map",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		dict := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &dict); err != nil {
			exit(commandName, logger, decoration, err)
		}
		key := args[1]
		fmt.Println(dict[key])
	},
}

func init() {
	rootCmd.AddCommand(getFromMapCmd)
}
