package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/explainer"
	"github.com/state-alchemists/zaruba/input"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/previousval"
	"github.com/state-alchemists/zaruba/runner"
	"github.com/state-alchemists/zaruba/strutil"
)

var pleaseEnvs []string
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
		decoration := cmdHelper.GetDecoration(*pleasePlainDecor)
		logger := output.NewConsoleLogger(decoration)
		csvRecordLogger := cmdHelper.GetCsvRecordLogger(filepath.Dir(pleaseFile))
		project, taskNames := getProjectAndTaskName(cmd, logger, decoration, args)
		prompter := input.NewPrompter(logger, decoration, project)
		explainer := explainer.NewExplainer(logger, decoration, project)
		isFallbackInteraction := false
		// no task provided
		if len(taskNames) == 0 && !*pleaseExplain {
			taskName := getTaskNameInteractivelyOrExit(cmd, logger, decoration, prompter)
			taskNames = []string{taskName}
			action := getActionOrExit(cmd, logger, decoration, prompter, taskName)
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
			initProjectOrExit(cmd, logger, decoration, project)
			explainOrExit(cmd, logger, decoration, explainer, taskNames)
			return
		}
		// handle "--previous"
		previousValueFile := ".previous.values.yaml"
		if *pleaseUsePreviousValues {
			loadPreviousValuesOrExit(cmd, logger, decoration, project, previousValueFile)
		}
		// handle "--interactive" flag
		if *pleaseInteractive {
			if !*pleaseUsePreviousValues {
				askProjectValueOrExit(cmd, logger, decoration, prompter)
			}
			askProjectEnvOrExit(cmd, logger, decoration, prompter, taskNames)
			askProjectValuesByTasksOrExit(cmd, logger, decoration, prompter, taskNames)
		}
		if isFallbackInteraction && !*pleaseTerminate {
			*pleaseTerminate = askAutoTerminateOrExit(cmd, logger, decoration, prompter, taskNames)
		}
		previousval.Save(project, previousValueFile)
		initProjectOrExit(cmd, logger, decoration, project)
		r, err := runner.NewRunner(logger, csvRecordLogger, project, taskNames, "10m", *pleaseTerminate, pleaseWait)
		if err != nil {
			showLastPleaseCommand(cmd, logger, decoration, taskNames)
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		if err := r.Run(); err != nil {
			showLastPleaseCommand(cmd, logger, decoration, taskNames)
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		showLastPleaseCommand(cmd, logger, decoration, taskNames)
	},
}

func showLastPleaseCommand(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, taskNames []string) {
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
	strUtil := strutil.NewStrUtil()
	// task names
	argTaskName := strings.Join(taskNames, " ")
	// value
	argValueList := []string{}
	for _, pleaseValue := range pleaseValues {
		argValueList = append(argValueList, fmt.Sprintf("-v %s", strUtil.EscapeShellValue(pleaseValue)))
	}
	argValue := strings.Join(argValueList, " ")
	// environment
	argEnvList := []string{}
	for _, pleaseEnv := range pleaseEnvs {
		argEnvList = append(argEnvList, fmt.Sprintf("-e %s", strUtil.EscapeShellValue(pleaseEnv)))
	}
	argEnv := strings.Join(argEnvList, " ")
	// terminate and wait
	argTerminate := ""
	argWait := ""
	if *pleaseTerminate {
		argTerminate += " -t"
		if pleaseWait != "0s" && pleaseWait != "" {
			argWait += fmt.Sprintf(" -w %s", pleaseWait)
		}
	}
	logger.Fprintf(os.Stderr, "%s%s %s %s %s%s%s%s\n", decoration.Yellow, commandName, argTaskName, argEnv, argValue, argTerminate, argWait, decoration.Normal)
}

func init() {
	// get current working directory
	dir, err := os.Getwd()
	if err != nil {
		dir = "."
	}
	// define defaultPleaseFile
	defaultPleaseFile := filepath.Join(dir, "index.zaruba.yaml")
	if _, err := os.Stat(defaultPleaseFile); os.IsNotExist(err) {
		defaultPleaseFile = "${ZARUBA_HOME}/core.zaruba.yaml"
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
	pleaseCmd.Flags().StringVarP(&pleaseFile, "file", "f", defaultPleaseFile, "project file")
	pleaseCmd.Flags().StringArrayVarP(&pleaseEnvs, "environment", "e", defaultEnv, "environments (e.g: '-e environment.env' or '-e KEY=VAL' or '-e {\"KEY\": \"VAL\"}' )")
	pleaseCmd.Flags().StringArrayVarP(&pleaseValues, "value", "v", defaultPleaseValues, "values (e.g: '-v value.yaml' or '-v key=val')")
	pleaseInteractive = pleaseCmd.Flags().BoolP("interactive", "i", false, "interactive mode")
	pleaseExplain = pleaseCmd.Flags().BoolP("explain", "x", false, "explain instead of execute")
	pleasePlainDecor = pleaseCmd.Flags().BoolP("nodecoration", "n", false, "no decoration")
	pleaseUsePreviousValues = pleaseCmd.Flags().BoolP("previous", "p", false, "load previous values")
	pleaseTerminate = pleaseCmd.Flags().BoolP("terminate", "t", false, "terminate after complete")
	pleaseCmd.Flags().StringVarP(&pleaseWait, "wait", "w", "0s", "termination waiting duration (e.g: '-w 5s'). Only take effect if -t or --terminate is set")
}

func getTaskNameInteractivelyOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter) (taskName string) {
	taskName, err := prompter.GetTaskName()
	if err != nil {
		cmdHelper.Exit(cmd, []string{}, logger, decoration, err)
	}
	return taskName
}

func getActionOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskName string) (action *input.Action) {
	action, err := prompter.GetAction(taskName)
	if err != nil {
		cmdHelper.Exit(cmd, []string{}, logger, decoration, err)
	}
	return action
}

func explainOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, explainer *explainer.Explainer, taskNames []string) {
	if err := explainer.Explain(taskNames...); err != nil {
		cmdHelper.Exit(cmd, []string{}, logger, decoration, err)
	}
}

func loadPreviousValuesOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, project *core.Project, previousValueFile string) {
	if err := previousval.Load(project, previousValueFile); err != nil {
		cmdHelper.Exit(cmd, []string{}, logger, decoration, err)
	}
}

func askProjectValuesByTasksOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskNames []string) {
	if err := prompter.SetProjectValuesByTask(taskNames); err != nil {
		cmdHelper.Exit(cmd, []string{}, logger, decoration, err)
	}
}

func askProjectEnvOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskNames []string) {
	if err := prompter.GetAdditionalEnv(taskNames); err != nil {
		cmdHelper.Exit(cmd, []string{}, logger, decoration, err)
	}
}

func askProjectValueOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter) {
	if err := prompter.GetAdditionalValue(); err != nil {
		cmdHelper.Exit(cmd, []string{}, logger, decoration, err)
	}
}

func askAutoTerminateOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskNames []string) (autoTerminate bool) {
	autoTerminate, err := prompter.GetAutoTerminate(taskNames)
	if err != nil {
		cmdHelper.Exit(cmd, []string{}, logger, decoration, err)
	}
	return autoTerminate
}

func initProjectOrExit(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, project *core.Project) {
	if err := project.Init(); err != nil {
		cmdHelper.Exit(cmd, []string{}, logger, decoration, err)
	}
}

func getProjectAndTaskName(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, args []string) (project *core.Project, taskNames []string) {
	project, err := core.NewProject(pleaseFile, decoration)
	if err != nil {
		cmdHelper.Exit(cmd, args, logger, decoration, err)
	}
	for _, env := range pleaseEnvs {
		if err = project.AddGlobalEnv(env); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	}
	// process values from flag
	for _, value := range pleaseValues {
		if err = project.AddValue(value); err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
	}
	taskNames = []string{}
	//  distinguish taskNames and additional values
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			if err = project.AddValue(arg); err != nil {
				cmdHelper.Exit(cmd, args, logger, decoration, err)
			}
			continue
		}
		taskNames = append(taskNames, arg)
	}
	return project, taskNames
}
