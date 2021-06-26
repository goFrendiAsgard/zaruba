package cmd

import (
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
	Short: "Declarative Task Runner Framework",
	Long: `💀 Declarative Task Runner Framework

Zaruba help you execute tasks faster and easier.
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
		logger.DPrintfError("Do you mean %s%szaruba please%s?\n", decoration.Bold, decoration.Yellow, decoration.Normal)
		os.Exit(1)
	}
}

func showErrorAndExit(logger output.Logger, decoration *output.Decoration, err error) {
	if err != nil {
		logger.Fprintf(os.Stderr, "%s %s%s%s%s\n", decoration.Error, decoration.Bold, decoration.Red, err.Error(), decoration.Normal)
		os.Exit(1)
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
