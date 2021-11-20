package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var strReplaceCmd = &cobra.Command{
	Use:   "replace <string> <replacementMap>",
	Short: "Replace string by replacementMap",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		text, mapString := args[0], args[1]
		util := core.NewCoreUtil()
		replacementMap, err := util.Json.Map.GetStringDict(mapString)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		result := util.Str.Replace(text, replacementMap)
		fmt.Println(result)
	},
}
