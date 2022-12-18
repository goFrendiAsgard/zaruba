package numcmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var randomCmd = &cobra.Command{
	Use:   "random [intStart] <intEnd>",
	Short: "Print a single random number",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		intStart, intEnd := 0, 0
		var err error
		switch len(args) {
		case 1:
			intEnd, err = strconv.Atoi(args[0])
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		case 2:
			intStart, err = strconv.Atoi(args[0])
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			intEnd, err = strconv.Atoi(args[1])
			if err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
		}
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Intn(intEnd-intStart) + intStart)
	},
}
