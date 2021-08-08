package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var getPrefixedEnvCmd = &cobra.Command{
	Use:   "getPrefixedEnv <prefix> <var>",
	Short: "Get prefixed environment value",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		prefix := args[0]
		key := args[1]
		if val := os.Getenv(fmt.Sprintf("%s_%s", prefix, key)); val != "" {
			fmt.Println(val)
			return
		}
		fmt.Println(os.Getenv(key))
	},
}

func init() {
	rootCmd.AddCommand(getPrefixedEnvCmd)
}
