package cmd

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var strSubmatchCmd = &cobra.Command{
	Use:   "strSubmatch <string> <pattern>",
	Short: "Return submatch of string based on pattern",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		text := args[0]
		rex, err := regexp.Compile(args[1])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		result := rex.FindStringSubmatch(text)
		resultB, err := json.Marshal(result)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}

func init() {
	rootCmd.AddCommand(strSubmatchCmd)
}
