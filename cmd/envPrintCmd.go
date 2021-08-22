package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var envPrintCmd = &cobra.Command{
	Use:   "print <envMap> [prefix]",
	Short: "Print environment",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		envMapRaw := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[0]), &envMapRaw); err != nil {
			exit(cmd, logger, decoration, err)
		}
		envMap := convertToMapString(envMapRaw)
		// cascade prefix
		if len(args) > 1 {
			prefix := args[1]
			envMap = envCascadePrefix(envMap, prefix)
		}
		result, err := godotenv.Marshal(envMap)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(result)
	},
}
