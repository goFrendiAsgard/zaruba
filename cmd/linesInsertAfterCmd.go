package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/utility"
)

var linesInsertAfterCmd = &cobra.Command{
	Use:   "insertAfter <lines> <index> <newLine>",
	Short: "Replace lines[index] with replacements",
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
		util := utility.NewUtil()
		result, err := util.Str.ReplaceLineAtIndex(lines, index, append([]string{lines[index]}, newLines...))
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
