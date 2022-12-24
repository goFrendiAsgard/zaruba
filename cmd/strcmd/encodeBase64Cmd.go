package strcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var encodeBase64Cmd = &cobra.Command{
	Use:   "encodeBase64 <string>",
	Short: "Encode base64 a string",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		text := args[0]
		util := dsl.NewDSLUtil()
		fmt.Println(util.Str.EncodeBase64(text))
	},
}
