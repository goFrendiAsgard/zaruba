package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/output"
)

var writeLinesCmd = &cobra.Command{
	Use:   "writeLines <fileName> <list>",
	Short: "Write list to file",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		fileName := args[0]
		list := []string{}
		if err := json.Unmarshal([]byte(args[1]), &list); err != nil {
			exit(commandName, logger, decoration, err)
		}
		if err := file.WriteLines(fileName, list, 0755); err != nil {
			exit(commandName, logger, decoration, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(writeLinesCmd)
}
