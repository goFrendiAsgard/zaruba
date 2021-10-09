package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/utility"
)

var strReplaceCmd = &cobra.Command{
	Use:   "replace <string> <replacementMap>",
	Short: "Replace string by replacementMap",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		text := args[0]
		rawReplacementMap := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args[1]), &rawReplacementMap); err != nil {
			exit(cmd, logger, decoration, err)
		}
		replacementMap := convertToMapString(rawReplacementMap)
		util := utility.NewUtil()
		result := util.Str.Replace(text, replacementMap)
		fmt.Println(result)
	},
}
