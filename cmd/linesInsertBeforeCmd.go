package cmd

import (
	"encoding/json"
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
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		lines := []string{}
		if err := json.Unmarshal([]byte(args[0]), &lines); err != nil {
			exit(cmd, logger, decoration, err)
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		newLines := []string{}
		if err := json.Unmarshal([]byte(args[2]), &newLines); err != nil {
			newLines = []string{args[2]}
		}
		util := core.NewCoreUtil()
		result, err := util.Str.ReplaceLineAtIndex(lines, index, append(newLines, lines[index]))
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		resultB, err := json.Marshal(result)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
