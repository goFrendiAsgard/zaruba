package timecmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var toTimestampExample = `
> zaruba time toTimestamp '2023-01-01T13:01:01.345Z'
1672578061

> zaruba time toTimestamp '2023-01-01T20:01:01+07:00'
1672578061
`

var toTimestamp = &cobra.Command{
	Use:     "toTimestamp <RFC-3339>",
	Short:   "Encode a string with time algorithm",
	Example: toTimestampExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		rfc3339 := args[0]
		t, err := time.Parse(time.RFC3339, rfc3339)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		fmt.Println(t.Unix())
	},
}
