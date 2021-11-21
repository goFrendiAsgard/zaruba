package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var linesInsertBeforeCmd = &cobra.Command{
	Use:   "insertBefore <lines> <index> <newLine>",
	Short: "Insert newLine before lines[index]",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		util := core.NewCoreUtil()
		jsonLines, jsonReplacements := args[0], args[2]
		index, err := strconv.Atoi(args[1])
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		newJsonLines, err := util.Json.List.InsertLineBeforeIndex(jsonLines, index, jsonReplacements)
		if err != nil {
			exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(newJsonLines)
	},
}
