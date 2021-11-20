package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var linesFillCmd = &cobra.Command{
	Use:   "fill <lines> <patterns> <suplements>",
	Short: "Insert suplements to lines if patterns is not found",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 3)
		lines := []string{}
		if err := json.Unmarshal([]byte(args[0]), &lines); err != nil {
			exit(cmd, logger, decoration, err)
		}
		patterns := []string{}
		if err := json.Unmarshal([]byte(args[1]), &patterns); err != nil {
			exit(cmd, logger, decoration, err)
		}
		suplements := []string{}
		if err := json.Unmarshal([]byte(args[2]), &suplements); err != nil {
			suplements = []string{args[2]}
		}
		util := core.NewCoreUtil()
		result, err := util.Str.CompleteLines(lines, patterns, suplements)
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
