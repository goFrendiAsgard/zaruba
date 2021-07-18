package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

var getServiceNameCmd = &cobra.Command{
	Use:   "getServiceName <string>",
	Short: "Get default service name based on location or image name",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 1)
		serviceName, err := config.GetDefaultServiceName(args[0])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(serviceName)
	},
}

func init() {
	rootCmd.AddCommand(getServiceNameCmd)
}
