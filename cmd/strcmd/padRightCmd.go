package strcmd

import (
	"fmt"
	"strconv"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var padRightCmd = &cobra.Command{
	Use:   "padRight <string> <length> [char]",
	Short: "Fill from right",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		text := args[0]
		length, err := strconv.Atoi(args[1])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		pad := " "
		if len(args) > 2 {
			pad = args[2]
		}
		util := dsl.NewDSLUtil()
		fmt.Println(util.Str.PadRight(text, length, pad))
	},
}
