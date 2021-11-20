package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var mapMergeCmd = &cobra.Command{
	Use:   "merge <jsonMap> <otherJsonMaps...>",
	Short: "Merge JSON maps, in case of duplicate keys, the first ocurrance is going to be used",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		mapStrings := args
		util := core.NewCoreUtil()
		mergedMapString, err := util.Json.Map.Merge(mapStrings...)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(mergedMapString)
	},
}
