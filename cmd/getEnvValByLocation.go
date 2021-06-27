package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var getEnvValByLocationCmd = &cobra.Command{
	Use:   "getEnvValByLocation",
	Short: "Get environment value by location",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 2 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for getEnvByLocation"))
		}
		envMap, err := util.GetEnvByLocation(args[0])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		val, envExist := envMap[args[1]]
		if !envExist {
			showErrorAndExit(logger, decoration, fmt.Errorf("environment not found in %s: %s", args[0], args[1]))
		}
		fmt.Println(val)
	},
}

func init() {
	rootCmd.AddCommand(getEnvValByLocationCmd)
}
