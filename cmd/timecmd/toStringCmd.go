package timecmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var toStringExample = `
> zaruba time toString 1672578061
2023-01-01T20:01:01+07:00
`

var toStringCmd = &cobra.Command{
	Use:     "toString <timestamp>",
	Short:   "Change timestamp into RFC3339 string",
	Example: toStringExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		timestamp, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(time.Unix(timestamp, 0).Format(time.RFC3339))
	},
}
