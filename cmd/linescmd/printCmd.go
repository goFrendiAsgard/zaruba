package linescmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var printCmd = &cobra.Command{
	Use:     "print <lines> [fileName]",
	Short:   "Print lines as multiline string",
	Aliases: []string{"write", "join"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		jsonString := args[0]
		fileName := ""
		util := dsl.NewDSLUtil()
		if len(args) > 1 {
			fileName = args[1]
			if err := util.File.WriteLines(fileName, jsonString, 0755); err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			return
		}
		stringList, err := util.Json.ToStringList(jsonString)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(strings.Join(stringList, "\n"))
	},
}
