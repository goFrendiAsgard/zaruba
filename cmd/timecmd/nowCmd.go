package timecmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Get current time as RFC3339",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 0)
		fmt.Println(time.Now().Format(time.RFC3339))
	},
}
