package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zaruba",
	Short: "Declarative task runner framework and CLI utilities",
	Long: `Declarative task runner framework and CLI utilities

Zaruba helps you declare/execute tasks to manage your projects.
Try:
  zaruba please`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		logger.Println(err)
		logger.DPrintfError("To run a task you need to type %s%szaruba please <task-name>%s\n", decoration.Bold, decoration.Yellow, decoration.Normal)
		os.Exit(1)
	}
}

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

func getProject(logger output.Logger, decoration *output.Decoration, csvRecordLogger *output.CSVRecordLogger, pleaseFile string) (project *config.Project, err error) {
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
	return config.NewProject(logger, csvRecordLogger, decoration, pleaseFile, defaultIncludes)
}
