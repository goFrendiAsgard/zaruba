package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var getEnvKeysByLocationCmd = &cobra.Command{
	Use:   "getEnvKeysByLocation",
	Short: "Get environment keys by location",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for getEnvByLocation"))
		}
		envMap, err := util.GetEnvByLocation(args[0])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		keys := []string{}
		for key, _ := range envMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println(key)
		}
	},
}

func init() {
	rootCmd.AddCommand(getEnvKeysByLocationCmd)
}
