package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

var strGetIndentationCmd = &cobra.Command{
	Use:   "strGetIndentation <string> <level>",
	Short: "Get indentation based on level",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		text := args[0]
		level, err := strconv.Atoi(args[1])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		result, err := str.GetSingleIndentation(text, level)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(strGetIndentationCmd)
}
