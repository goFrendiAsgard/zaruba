package base32cmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var encodeCmd = &cobra.Command{
	Use:   "encode <string>",
	Short: "Encode a string with base32 algorithm",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		util := dsl.NewDSLUtil()
		fmt.Println(util.Str.EncodeBase32(text))
	},
}
