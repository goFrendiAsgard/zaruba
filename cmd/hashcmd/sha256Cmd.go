package hashcmd

import (
	"crypto/sha256"
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var sha256Cmd = &cobra.Command{
	Use:   "sha256 <str>",
	Short: "Has str with sha256",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		str := args[0]
		hashedStr := fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
		fmt.Println(hashedStr)
	},
}
