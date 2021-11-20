package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var linesSubmatchCmd = &cobra.Command{
	Use:   "submatch <list> <patterns>",
	Short: "Return submatch matching the pattern",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		list := []string{}
		err := json.Unmarshal([]byte(args[0]), &list)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		patterns := []string{}
		if err := json.Unmarshal([]byte(args[1]), &patterns); err != nil {
			exit(cmd, logger, decoration, err)
		}
		util := core.NewCoreUtil()
		index, submatch, err := util.Str.GetLineSubmatch(list, patterns)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		if index == -1 {
			exit(cmd, logger, decoration, fmt.Errorf("no line match %#v", patterns))
		}
		resultB, err := json.Marshal(submatch)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
