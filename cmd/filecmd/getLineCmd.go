package filecmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var getLineExample = `
Getting line index that match the last element of the pattern
    > zaruba file read myFile.txt
    a
    a
    b
    c
    d
    e
    > zaruba file getLine myFile.txt 0
    a
    > zaruba file getLine myFile.txt 2
    b
    > zaruba file getLine myFile.txt -1
    e
`
var getLineCmd = &cobra.Command{
	Use:     "getLine <strFileName> <index>",
	Short:   "Return desired line of a file content",
	Example: getLineExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		util := dsl.NewDSLUtil()
		strFileName, indexStr := args[0], args[1]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		jsonLines, err := util.File.ReadLines(strFileName)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		stringList, err := util.Json.ToStringList(jsonLines)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if index < 0 {
			index = len(stringList) + index
		}
		fmt.Println(stringList[index])
	},
}
