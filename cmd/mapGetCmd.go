package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var mapGetCmd = &cobra.Command{
	Use:   "get <map> <key>",
	Short: "Get value from JSON map",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		dict := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &dict); err != nil {
			exit(cmd, logger, decoration, err)
		}
		key := args[1]
		printInterface(dict[key])
	},
}
