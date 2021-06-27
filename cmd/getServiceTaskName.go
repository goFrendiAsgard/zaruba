package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var getServiceTaskNameCmd = &cobra.Command{
	Use:   "getServiceTaskName <serviceName>",
	Short: "Get task name of a service",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for getServiceTaskName"))
		}
		taskName := util.GetServiceTaskName(args[0])
		fmt.Println(taskName)
	},
}

func init() {
	rootCmd.AddCommand(getServiceTaskNameCmd)
}
