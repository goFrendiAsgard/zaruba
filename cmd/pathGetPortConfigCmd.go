package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/env"
	"github.com/state-alchemists/zaruba/output"
)

var pathGetPortConfigCmd = &cobra.Command{
	Use:   "getPortConfig <location>",
	Short: "Return string representing default config.ports",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		envMap, err := env.GetEnvByLocation(args[0])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		ports := []string{}
		for key, val := range envMap {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			if intVal > 1000 {
				ports = append(ports, fmt.Sprintf("{{ .GetEnv \"%s\" }}", key))
			}
		}
		resultB, err := json.Marshal(ports)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
