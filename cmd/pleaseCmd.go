package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/explainer"
	"github.com/state-alchemists/zaruba/input"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/previousval"
	"github.com/state-alchemists/zaruba/runner"
)

var pleaseEnv []string
var pleaseValues []string
var pleaseFile string
var pleaseInteractive *bool
var pleaseUsePreviousValues *bool
var pleaseTerminate *bool
var pleaseExplain *bool
var pleasePlainDecor *bool
var pleaseWait string

// pleaseCmd represents the please command
var pleaseCmd = &cobra.Command{
	Use:     "please [taskName...]",
	Short:   "Run Task(s)",
	Aliases: []string{"run", "do", "execute", "exec", "perform", "invoke"},
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := getDecoration(*pleasePlainDecor)
		logger := output.NewConsoleLogger(decoration)
		csvRecordLogger := getCsvRecordLogger(filepath.Dir(pleaseFile))
		project, taskNames := getProjectAndTaskName(commandName, logger, decoration, csvRecordLogger, args)
		prompter := input.NewPrompter(logger, decoration, project)
		explainer := explainer.NewExplainer(logger, decoration, project)
		isFallbackInteraction := false
		// no task provided
		if len(taskNames) == 0 && !*pleaseExplain {
			taskName := getTaskNameInteractivelyOrExit(commandName, logger, decoration, prompter)
			taskNames = []string{taskName}
			action := getActionOrExit(commandName, logger, decoration, prompter, taskName)
			if action.Explain {
				*pleaseExplain = true
			}
			if action.RunInteractive {
				*pleaseInteractive = true
			}
			isFallbackInteraction = true
		}
		// handle "--explain"
		if *pleaseExplain {
			initProjectOrExit(commandName, logger, decoration, project)
			explainOrExit(commandName, logger, decoration, explainer, taskNames)
			return
		}
		// handle "--previous"
		previousValueFile := ".previous.values.yaml"
		if *pleaseUsePreviousValues {
			loadPreviousValuesOrExit(commandName, logger, decoration, project, previousValueFile)
		}
		// handle "--interactive" flag
		if *pleaseInteractive {
			if !*pleaseUsePreviousValues {
				askProjectValueOrExit(commandName, logger, decoration, prompter)
			}
			askProjectEnvOrExit(commandName, logger, decoration, prompter, taskNames)
			askProjectValuesByTasksOrExit(commandName, logger, decoration, prompter, taskNames)
		}
		if isFallbackInteraction && !*pleaseTerminate {
			*pleaseTerminate = askAutoTerminateOrExit(commandName, logger, decoration, prompter, taskNames)
		}
		previousval.Save(project, previousValueFile)
		initProjectOrExit(commandName, logger, decoration, project)
		r, err := runner.NewRunner(logger, decoration, project, taskNames, "5m", *pleaseTerminate, pleaseWait)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		if err := r.Run(); err != nil {
			exit(commandName, logger, decoration, err)
		}
		logger.DPrintf("%sLast command:%s %s\n", decoration.Yellow, decoration.Normal, explainer.GetZarubaCommand(taskNames, *pleaseTerminate, pleaseWait))
	},
}

func init() {
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
	pleaseCmd.Flags().StringVarP(&pleaseFile, "file", "f", defaultPleaseFile, "project script file")
	pleaseCmd.Flags().StringArrayVarP(&pleaseEnv, "environment", "e", defaultEnv, "environment file or pairs (e.g: '-e environment.env' or '-e key=val')")
	pleaseCmd.Flags().StringArrayVarP(&pleaseValues, "value", "v", defaultPleaseValues, "yaml file or pairs (e.g: '-v value.yaml' or '-v key=val')")
	pleaseInteractive = pleaseCmd.Flags().BoolP("interactive", "i", false, "if set, you will be able to input values interactively")
	pleaseExplain = pleaseCmd.Flags().BoolP("explain", "x", false, "if set, the tasks will be explained instead of executed")
	pleasePlainDecor = pleaseCmd.Flags().BoolP("nodecoration", "n", false, "if set, there will be no decoration")
	pleaseUsePreviousValues = pleaseCmd.Flags().BoolP("previous", "p", false, "if set, previous values will be loaded")
	pleaseTerminate = pleaseCmd.Flags().BoolP("terminate", "t", false, "if set, tasks will be terminated after complete")
	pleaseCmd.Flags().StringVarP(&pleaseWait, "wait", "w", "0s", "how long zaruba should wait before terminating tasks (e.g: '-w 5s'). Only take effect if -t or --terminate is set")
}

func getTaskNameInteractivelyOrExit(commandName string, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter) (taskName string) {
	taskName, err := prompter.GetTaskName()
	if err != nil {
		exit(commandName, logger, decoration, err)
	}
	return taskName
}

func getActionOrExit(commandName string, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskName string) (action *input.Action) {
	action, err := prompter.GetAction(taskName)
	if err != nil {
		exit(commandName, logger, decoration, err)
	}
	return action
}

func explainOrExit(commandName string, logger *output.ConsoleLogger, decoration *output.Decoration, explainer *explainer.Explainer, taskNames []string) {
	if err := explainer.Explain(taskNames...); err != nil {
		exit(commandName, logger, decoration, err)
	}
}

func loadPreviousValuesOrExit(commandName string, logger *output.ConsoleLogger, decoration *output.Decoration, project *config.Project, previousValueFile string) {
	if err := previousval.Load(project, previousValueFile); err != nil {
		exit(commandName, logger, decoration, err)
	}
}

func askProjectValuesByTasksOrExit(commandName string, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskNames []string) {
	if err := prompter.SetProjectValuesByTask(taskNames); err != nil {
		exit(commandName, logger, decoration, err)
	}
}

func askProjectEnvOrExit(commandName string, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskNames []string) {
	if err := prompter.GetAdditionalEnv(taskNames); err != nil {
		exit(commandName, logger, decoration, err)
	}
}

func askProjectValueOrExit(commandName string, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter) {
	if err := prompter.GetAdditionalValue(); err != nil {
		exit(commandName, logger, decoration, err)
	}
}

func askAutoTerminateOrExit(commandName string, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskNames []string) (autoTerminate bool) {
	autoTerminate, err := prompter.GetAutoTerminate(taskNames)
	if err != nil {
		exit(commandName, logger, decoration, err)
	}
	return autoTerminate
}

func initProjectOrExit(commandName string, logger output.Logger, decoration *output.Decoration, project *config.Project) {
	if err := project.Init(); err != nil {
		exit(commandName, logger, decoration, err)
	}
}

func getProjectAndTaskName(commandName string, logger output.Logger, decoration *output.Decoration, csvRecordLogger *output.CSVRecordLogger, args []string) (project *config.Project, taskNames []string) {
	project, err := getProject(logger, decoration, csvRecordLogger, pleaseFile)
	if err != nil {
		exit(commandName, logger, decoration, err)
	}
	for _, env := range pleaseEnv {
		if err = project.AddGlobalEnv(env); err != nil {
			exit(commandName, logger, decoration, err)
		}
	}
	// process values from flag
	for _, value := range pleaseValues {
		if err = project.AddValue(value); err != nil {
			exit(commandName, logger, decoration, err)
		}
	}
	taskNames = []string{}
	//  distinguish taskNames and additional values
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			if err = project.AddValue(arg); err != nil {
				exit(commandName, logger, decoration, err)
			}
			continue
		}
		taskNames = append(taskNames, arg)
	}
	return project, taskNames
}
