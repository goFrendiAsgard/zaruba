package cmd

import (
	"encoding/json"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var envWriteCmd = &cobra.Command{
	Use:   "write <fileName> <envMap> [prefix]",
	Short: "Write envMap to file",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		fileName := args[0]
		envMapRaw := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[1]), &envMapRaw); err != nil {
			exit(commandName, logger, decoration, err)
		}
		envMap := convertToMapString(envMapRaw)
		// cascade prefix
		if len(args) > 2 {
			prefix := args[2]
			envMap = envCascadePrefix(envMap, prefix)
		}
		if err := godotenv.Write(envMap, fileName); err != nil {
			exit(commandName, logger, decoration, err)
		}
	},
}
