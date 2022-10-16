package numcmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var rangeCmd = &cobra.Command{
	Use:   "range [start] <end> [step]",
	Short: "Print numbers sequentially",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		start, end, step := 0, 0, 1
		var err error
		switch len(args) {
		case 1:
			end, err = strconv.Atoi(args[0])
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		case 2:
			start, err = strconv.Atoi(args[0])
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			end, err = strconv.Atoi(args[1])
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		case 3:
			start, err = strconv.Atoi(args[0])
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			end, err = strconv.Atoi(args[1])
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			step, err = strconv.Atoi(args[2])
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		}
		for i := start; i < end; i += step {
			fmt.Println(i)
		}
	},
}
