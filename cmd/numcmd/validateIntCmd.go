package numcmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var validateIntCmd = &cobra.Command{
	Use:   "validateInt <value>",
	Short: "Check whether value is valid int or not",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println(0)
			return
		}
		fmt.Println(1)
	},
}
