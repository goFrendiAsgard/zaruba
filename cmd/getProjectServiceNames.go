package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var getProjectServiceNamesCmd = &cobra.Command{
	Use:   "getProjectServiceNames <projectFile>",
	Short: "Get project's service names",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for getProjectServiceNames"))
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
		if err = project.Init(); err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		for _, serviceName := range util.GetProjectServiceNames(project) {
			fmt.Println(serviceName)
		}
	},
}

func init() {
	rootCmd.AddCommand(getProjectServiceNamesCmd)
}
