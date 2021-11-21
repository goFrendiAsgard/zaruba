package mapcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var mergeCmd = &cobra.Command{
	Use:   "merge <jsonMap> <otherJsonMaps...>",
	Short: "Merge JSON maps, in case of duplicate keys, the first ocurrance is going to be used",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		mapStrings := args
		util := core.NewCoreUtil()
		mergedMapString, err := util.Json.Map.Merge(mapStrings...)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(mergedMapString)
	},
}
