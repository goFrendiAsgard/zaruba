package numcmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var rangeCmd = &cobra.Command{
	Use:   "range [intStart] <intEnd> [intStep]",
	Short: "Print numbers sequentially",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		intStart, intEnd, intStep := 0, 0, 1
		var err error
		switch len(args) {
		case 1:
			intEnd, err = strconv.Atoi(args[0])
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
		case 2:
			intStart, err = strconv.Atoi(args[0])
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
			intEnd, err = strconv.Atoi(args[1])
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
		case 3:
			intStart, err = strconv.Atoi(args[0])
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
			intEnd, err = strconv.Atoi(args[1])
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
			intStep, err = strconv.Atoi(args[2])
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
		}
		for i := intStart; i < intEnd; i += intStep {
			fmt.Println(i)
		}
	},
}
