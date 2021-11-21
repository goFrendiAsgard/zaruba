package strcmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var repeatCmd = &cobra.Command{
	Use:   "repeat <string> <repetition>",
	Short: "Repeat string for repetition times",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		text := args[0]
		repetition, err := strconv.Atoi(args[1])
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		util := core.NewCoreUtil()
		fmt.Println(util.Str.Repeat(text, repetition))
	},
}
