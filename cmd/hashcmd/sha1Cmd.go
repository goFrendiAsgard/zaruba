package hashcmd

import (
	"crypto/sha1"
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var sha1Cmd = &cobra.Command{
	Use:   "sha1 <str>",
	Short: "Has str with sha1",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		str := args[0]
		hashedStr := fmt.Sprintf("%x", sha1.Sum([]byte(str)))
		fmt.Println(hashedStr)
	},
}
