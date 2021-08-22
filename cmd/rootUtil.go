package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

func exit(commandName string, logger output.Logger, decoration *output.Decoration, err error) {
	if err != nil {
		logger.Fprintf(os.Stderr, "%s %s%s%s: %s%s\n", decoration.Error, decoration.Bold, decoration.Red, commandName, err.Error(), decoration.Normal)
		os.Exit(1)
	}
}

func checkMinArgCount(commandName string, logger output.Logger, decoration *output.Decoration, args []string, minArgCount int) {
	if len(args) < minArgCount {
		err := fmt.Errorf("expecting %d arguments, get %#v", minArgCount, args)
		exit(commandName, logger, decoration, err)
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
