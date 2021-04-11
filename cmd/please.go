package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/explainer"
	"github.com/state-alchemists/zaruba/input"
	"github.com/state-alchemists/zaruba/monitor"
	"github.com/state-alchemists/zaruba/previousval"
	"github.com/state-alchemists/zaruba/runner"
)

var pleaseEnv []string
var pleaseValues []string
var pleaseFile string
var pleaseInteractive *bool
var pleaseUsePreviousValues *bool
var pleaseTerminate *bool
var pleaseWait string

// pleaseCmd represents the please command
var pleaseCmd = &cobra.Command{
	Use:     "please",
	Short:   "Ask Zaruba to do something for you",
	Long:    "💀 Ask Zaruba to do something for you",
	Aliases: []string{"run", "do", "execute", "exec", "perform", "invoke"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := monitor.NewDecoration()
		logger := monitor.NewConsoleLogger(decoration)
		project, taskNames := getProjectOrExit(logger, decoration, args)
		prompter := input.NewPrompter(logger, decoration, project)
		// no task provided
		if len(taskNames) == 0 {
			taskName, err := prompter.GetTaskName()
			if err != nil {
				showErrorAndExit(logger, decoration, err)
			}
			taskNames = []string{taskName}
			action, err := prompter.GetAction(taskName)
			if err != nil {
				showErrorAndExit(logger, decoration, err)
			}
			if action.Explain {
				initProjectOrExit(logger, decoration, project)
				exp := explainer.NewExplainer(logger, decoration, project)
				exp.Explain(taskName)
				return
			}
			if action.RunInteractive {
				*pleaseInteractive = true
			}
		}
		// handle "--previous"
		previousValueFile := ".previous.values.yaml"
		if *pleaseUsePreviousValues {
			if err := previousval.Load(project, previousValueFile); err != nil {
				showErrorAndExit(logger, decoration, err)
			}
		}
		// handle "--interactive" flag
		if *pleaseInteractive {
			if err := prompter.SetProjectValuesByTask(taskNames); err != nil {
				showErrorAndExit(logger, decoration, err)
			}
		}
		previousval.Save(project, previousValueFile)
		initProjectOrExit(logger, decoration, project)
		r, err := runner.NewRunner(logger, decoration, project, taskNames, time.Minute*5)
		if *pleaseTerminate {
			r.SetTerminationDelay(pleaseWait)
		}
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		if err := r.Run(); err != nil {
			showErrorAndExit(logger, decoration, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pleaseCmd)
	// get current working directory
	dir, err := os.Getwd()
	if err != nil {
		dir = "."
	}
	// define defaultPleaseFile
	defaultPleaseFile := filepath.Join(dir, "main.zaruba.yaml")
	if _, err := os.Stat(defaultPleaseFile); os.IsNotExist(err) {
		defaultPleaseFile = "${ZARUBA_HOME}/scripts/core.zaruba.yaml"
	}
	// define defaultPleaseValues
	defaultPleaseValues := []string{}
	defaultValuesFile := filepath.Join(dir, "default.values.yaml")
	if _, err := os.Stat(defaultValuesFile); !os.IsNotExist(err) {
		defaultPleaseValues = append(defaultPleaseValues, defaultValuesFile)
	}
	// define defaultEnvFile
	defaultEnv := []string{}
	defaultEnvFile := filepath.Join(dir, ".env")
	if _, err := os.Stat(defaultEnvFile); !os.IsNotExist(err) {
		defaultEnv = append(defaultEnv, defaultEnvFile)
	}
	// register flags
	pleaseCmd.Flags().StringVarP(&pleaseFile, "file", "f", defaultPleaseFile, "main zaruba script file")
	pleaseCmd.Flags().StringArrayVarP(&pleaseEnv, "environment", "e", defaultEnv, "environment file or pairs (e.g: '-e environment.env' or '-e key=val')")
	pleaseCmd.Flags().StringArrayVarP(&pleaseValues, "value", "v", defaultPleaseValues, "yaml file or pairs (e.g: '-v value.yaml' or '-v key=val')")
	pleaseInteractive = pleaseCmd.Flags().BoolP("interactive", "i", false, "if set, you will be able to input values interactively")
	pleaseUsePreviousValues = pleaseCmd.Flags().BoolP("previous", "p", false, "if set, previous values will be loaded")
	pleaseTerminate = pleaseCmd.Flags().BoolP("terminate", "t", false, "if set, tasks will be terminated after complete")
	pleaseCmd.Flags().StringVarP(&pleaseWait, "wait", "w", "0s", "how long zaruba should wait before terminating tasks (e.g: '-w 5s'). Only take effect if -t or --terminate is set")
}

func initProjectOrExit(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project) {
	if err := project.Init(); err != nil {
		showErrorAndExit(logger, decoration, err)
	}
}

func getProjectOrExit(logger monitor.Logger, decoration *monitor.Decoration, args []string) (project *config.Project, taskNames []string) {
	project, taskNames, err := getProject(logger, decoration, args)
	if err != nil {
		showErrorAndExit(logger, decoration, err)
	}
	return project, taskNames
}

func getProject(logger monitor.Logger, decoration *monitor.Decoration, args []string) (project *config.Project, taskNames []string, err error) {
	taskNames = []string{}
	dir := os.ExpandEnv(filepath.Dir(pleaseFile))
	logFile := filepath.Join(dir, "log.zaruba.csv")
	csvLogger := monitor.NewCSVLogWriter(logFile)
	project, err = config.NewProject(logger, csvLogger, decoration, pleaseFile)
	if err != nil {
		return project, taskNames, err
	}
	// process globalEnv
	for _, env := range pleaseEnv {
		if err = project.AddGlobalEnv(env); err != nil {
			return project, taskNames, err
		}
	}
	// process values from flag
	for _, value := range pleaseValues {
		if err = project.AddValue(value); err != nil {
			return project, taskNames, err
		}
	}
	//  distinguish taskNames and additional values
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			if err = project.AddValue(arg); err != nil {
				return project, taskNames, err
			}
			continue
		}
		taskNames = append(taskNames, arg)
	}
	return project, taskNames, err
}

func showErrorAndExit(logger monitor.Logger, decoration *monitor.Decoration, err error) {
	if err != nil {
		logger.DPrintfError("%s%s%s%s\n", decoration.Bold, decoration.Red, err.Error(), decoration.Normal)
		os.Exit(1)
	}
}
