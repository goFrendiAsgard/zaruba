package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var linesIndentCmd = &cobra.Command{
	Use:   "indent <list> <indentation>",
	Short: "indent lines",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		lines := []string{}
		if err := json.Unmarshal([]byte(args[0]), &lines); err != nil {
			exit(cmd, logger, decoration, err)
		}
		indentation := args[1]
		for index, line := range lines {
			if strings.Trim(line, " ") != "" {
				lines[index] = indentation + line
			}
		}
		resultB, err := json.Marshal(lines)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
