package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/env"
	"github.com/state-alchemists/zaruba/output"
)

var pathGetEnvCmd = &cobra.Command{
	Use:   "getEnv <location>",
	Short: "Return JSON string map containing environment variables defined on location",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		result, err := env.GetEnvByLocation(args[0])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		resultB, err := json.Marshal(result)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
