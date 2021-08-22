package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var envGetCmd = &cobra.Command{
	Use:   "get [prefix]",
	Short: "Get envmap from currently loaded environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 0)
		// get envMap
		envMap := map[string]string{}
		for _, pair := range os.Environ() {
			pairParts := strings.SplitN(pair, "=", 2)
			key, val := pairParts[0], pairParts[1]
			envMap[key] = val
		}
		// cascade prefix
		if len(args) > 0 {
			prefix := args[0]
			envMap = envCascadePrefix(envMap, prefix)
		}
		// return the map
		resultB, err := json.Marshal(envMap)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
