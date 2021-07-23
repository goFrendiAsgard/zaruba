package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/env"
	"github.com/state-alchemists/zaruba/output"
)

var getPortConfigByLocationCmd = &cobra.Command{
	Use:   "getPortConfigByLocation <location>",
	Short: "Return string representing default config.ports",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		envMap, err := env.GetEnvByLocation(args[0])
		if err != nil {
			exit(commandName, logger, decoration, err)
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
		fmt.Println(strings.Join(ports, "\n"))
	},
}

func init() {
	rootCmd.AddCommand(getPortConfigByLocationCmd)
}
