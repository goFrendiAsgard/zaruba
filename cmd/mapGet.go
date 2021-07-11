package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var mapGetCmd = &cobra.Command{
	Use:   "mapGet <map> <key>",
	Short: "Get map[key]",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		dict := map[string]string{}
		if err := json.Unmarshal([]byte(args[0]), &dict); err != nil {
			exit(commandName, logger, decoration, err)
		}
		key := args[1]
		fmt.Println(dict[key])
	},
}

func init() {
	rootCmd.AddCommand(mapGetCmd)
}
