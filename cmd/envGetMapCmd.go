package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var envGetMapCmd = &cobra.Command{
	Use:   "getMap [prefix]",
	Short: "Get environment map",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		// get envMap
		envMap := map[string]string{}
		for _, pair := range os.Environ() {
			pairParts := strings.SplitN(pair, "=", 2)
			key := pairParts[0]
			val := pairParts[1]
			envMap[key] = val
		}
		// cascade
		if len(args) > 0 {
			prefix := args[1]
			for key := range envMap {
				prefixedKey := fmt.Sprintf("%s_%s", prefix, key)
				if prefixedVal, exist := envMap[prefixedKey]; exist {
					envMap[key] = prefixedVal
				}
			}
		}
		// return the map
		resultB, err := json.Marshal(envMap)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
