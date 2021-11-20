package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
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
		logger.Fprintf(os.Stderr,
			"%s %s%s%s%s\n%s %s%s%s%s\n",
			decoration.Error, decoration.Bold, decoration.Red, commandName, decoration.Normal,
			decoration.Error, decoration.Bold, decoration.Red, err.Error(), decoration.Normal)
		os.Exit(1)
	}
}

func checkMinArgCount(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, args []string, minArgCount int) {
	if len(args) < minArgCount {
		usage := cmd.UsageString()
		jsonB, _ := json.Marshal(args)
		err := fmt.Errorf("expecting %d arguments, get %d: %s\n%s", minArgCount, len(args), string(jsonB), usage)
		exit(cmd, logger, decoration, err)
	}
}

func getDecoration(plainDecor bool) (decoration *output.Decoration) {
	if plainDecor {
		return output.NewPlainDecoration()
	}
	return output.NewDefaultDecoration()
}

func getCsvRecordLogger(projectDir string) (csvRecordLogger *output.CSVRecordLogger) {
	logFile := filepath.Join(projectDir, "log.zaruba.csv")
	return output.NewCSVRecordLogger(logFile)
}

func getProject(decoration *output.Decoration, projectFile string) (project *core.Project, err error) {
	defaultIncludes := []string{"${ZARUBA_HOME}/core.zaruba.yaml"}
	for _, script := range strings.Split(os.Getenv("ZARUBA_SCRIPTS"), ":") {
		if script == "" {
			continue
		}
		defaultIncludes = append(defaultIncludes, script)
	}
	return core.NewCustomProject(decoration, projectFile, defaultIncludes)
}
