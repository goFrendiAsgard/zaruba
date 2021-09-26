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
		replacementMap := map[string]string{}
		if err := json.Unmarshal([]byte(args[1]), &replacementMap); err != nil {
			exit(cmd, logger, decoration, err)
		}
		util := utility.NewUtil()
		result := util.Str.Replace(text, replacementMap)
		fmt.Println(result)
	},
}
