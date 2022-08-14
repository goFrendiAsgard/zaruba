package strcmd

import (
	"fmt"
	"strconv"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var getIndentationCmd = &cobra.Command{
	Use:   "getIndentation <string> [level=1]",
	Short: "Get indentation of string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		level := 1
		if len(args) > 1 {
			if levelInput, err := strconv.Atoi(args[1]); err == nil {
				level = levelInput
			}
		}
		util := dsl.NewDSLUtil()
		result, err := util.Str.GetIndentation(text, level)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(result)
	},
}
