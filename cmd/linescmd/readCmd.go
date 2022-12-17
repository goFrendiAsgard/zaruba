package linescmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var readCmd = &cobra.Command{
	Use:   "read <strFileName>",
	Short: "Read lines from a file, return a jsonStrList",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		fileName := args[0]
		util := dsl.NewDSLUtil()
		jsonStrList, err := util.File.ReadLines(fileName)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(jsonStrList)
	},
}
