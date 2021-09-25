package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

func exit(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, err error) {
	if err != nil {
		nodeCmd := cmd
		commandName := ""
		for nodeCmd != nil {
			if commandName == "" {
				commandName = nodeCmd.Name()
			} else {
				commandName = fmt.Sprintf("%s %s", nodeCmd.Name(), commandName)
			}
			nodeCmd = nodeCmd.Parent()
		}
		usage := cmd.UsageString()
		logger.Fprintf(os.Stderr,
			"%s %s%s%s: %s\n%s%s\n", decoration.Error, decoration.Bold, decoration.Red, commandName, err.Error(), usage, decoration.Normal)
		os.Exit(1)
	}
}

func checkMinArgCount(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, args []string, minArgCount int) {
	if len(args) < minArgCount {
		jsonB, _ := json.Marshal(args)
		err := fmt.Errorf("expecting %d arguments, get %d: %s", minArgCount, len(args), string(jsonB))
		exit(cmd, logger, decoration, err)
	}
}

func getDecoration(plainDecor bool) (decoration *output.Decoration) {
	if plainDecor {
		return output.NewPlainDecoration()
	}
	return output.NewDecoration()
}

func getCsvRecordLogger(projectDir string) (csvRecordLogger *output.CSVRecordLogger) {
	logFile := filepath.Join(projectDir, "log.zaruba.csv")
	return output.NewCSVRecordLogger(logFile)
}

func getProject(decoration *output.Decoration, projectFile string) (project *config.Project, err error) {
	if os.Getenv("ZARUBA_HOME") == "" {
		executable, _ := os.Executable()
		os.Setenv("ZARUBA_HOME", filepath.Dir(executable))
	}
	defaultIncludes := []string{"${ZARUBA_HOME}/scripts/core.zaruba.yaml"}
	for _, script := range strings.Split(os.Getenv("ZARUBA_SCRIPTS"), ":") {
		if script == "" {
			continue
		}
		defaultIncludes = append(defaultIncludes, script)
	}
	return config.NewProject(decoration, projectFile, defaultIncludes)
}
