package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var getTaskLocationCmd = &cobra.Command{
	Use:   "getTaskLocation",
	Short: "Get task location by service name",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 2 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for getTaskLocation"))
		}
		projectFile, err := filepath.Abs(args[0])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		csvRecordLogger := getCsvRecordLogger(filepath.Dir(projectFile))
		project, err := getProject(logger, decoration, csvRecordLogger, projectFile)
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		serviceName, err := util.GetTaskLocation(project, args[1])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		fmt.Println(serviceName)
	},
}

func init() {
	rootCmd.AddCommand(getTaskLocationCmd)
}
