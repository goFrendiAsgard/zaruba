package hashcmd

import (
	"crypto/md5"
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var md5Cmd = &cobra.Command{
	Use:   "md5 <str>",
	Short: "Hash str with md5",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		str := args[0]
		hashedStr := fmt.Sprintf("%x", md5.Sum([]byte(str)))
		fmt.Println(hashedStr)
	},
}
