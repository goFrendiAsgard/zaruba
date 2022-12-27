package numcmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var addCmd = &cobra.Command{
	Use:   "add <floatNum1> <floatNum2> [floatNum3... floatNumN]",
	Short: "Add numbers",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		var result float64 = 0.0
		for _, arg := range args {
			element, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
			result += element
		}
		fmt.Println(result)
	},
}
