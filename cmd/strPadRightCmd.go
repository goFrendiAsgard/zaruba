package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var strPadRightCmd = &cobra.Command{
	Use:   "padRight <string> <length> [char]",
	Short: "fill from left",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		text := args[0]
		length, err := strconv.Atoi(args[1])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		pad := " "
		if len(args) > 2 {
			pad = args[2]
		}
		util := core.NewUtil()
		fmt.Println(util.Str.PadRight(text, length, pad))
	},
}
