package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/explainer"
	"github.com/state-alchemists/zaruba/helper"
	"github.com/state-alchemists/zaruba/input"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/previousval"
	"github.com/state-alchemists/zaruba/runner"
	"github.com/state-alchemists/zaruba/strutil"
)

var pleaseEnvs []string
var pleaseValues []string
var pleaseProjectFile string
var pleaseDecor string
var pleaseInteractive *bool
var pleaseUsePreviousValues *bool
var pleaseTerminate *bool
var pleaseShowLogTime *bool
var pleaseExplain *bool
var pleaseWait string

// pleaseCmd represents the please command
var pleaseCmd = &cobra.Command{
	Use:     "please [taskName...]",
	Short:   "Run Task(s)",
	Aliases: []string{"run", "do", "p"},
	Run: func(cmd *cobra.Command, args []string) {
		decoration := cmdHelper.GetDecoration(pleaseDecor)
		logger := output.NewConsoleLogger(decoration)
		csvRecordLogger := cmdHelper.GetCsvRecordLogger(filepath.Dir(pleaseProjectFile))
		project, taskNames := getProjectAndTaskName(cmd, logger, decoration, *pleaseShowLogTime, args)
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
		statusTimeIntervalStr := os.Getenv("ZARUBA_LOG_STATUS_TIME_INTERVAL")
		statusLineInterval, err := strconv.Atoi(os.Getenv("ZARUBA_LOG_STATUS_LINE_INTERVAL"))
		if err != nil {
			statusLineInterval = 0
		}
		r, err := runner.NewRunner(logger, csvRecordLogger, project, taskNames, statusTimeIntervalStr, statusLineInterval, *pleaseTerminate, pleaseWait)
		if err != nil {
			showLastPleaseCommand(cmd, logger, decoration)
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		if err := r.Run(); err != nil {
			showLastPleaseCommand(cmd, logger, decoration)
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		showLastPleaseCommand(cmd, logger, decoration)
	},
}

func init() {
	dsl.SetDefaultEnv()
	util := dsl.NewDSLUtil()
	// get current working directory
	dir, err := os.Getwd()
	if err != nil {
		dir = "."
	}
	zarubaEnv := os.Getenv("ZARUBA_ENV")
	defaultProjectFile := getDefaultProjectFile(dir)
	defaultValueFiles := getDefaultValueFiles(dir, zarubaEnv)
	defaultEnvFiles := getDefaultEnvFiles(dir, zarubaEnv)
	defaultDecoration := os.Getenv("ZARUBA_DECORATION")
	defaultShowLogTime := util.Bool.IsTrue(os.Getenv("ZARUBA_LOG_TIME"))
	// get parameters
	pleaseCmd.Flags().StringVarP(&pleaseProjectFile, "file", "f", defaultProjectFile, "project file")
	pleaseCmd.Flags().StringArrayVarP(&pleaseEnvs, "environment", "e", []string{}, "environments (e.g., '-e environment.env' or '-e KEY=VAL' or '-e {\"KEY\": \"VAL\"}' )")
	pleaseEnvs = append(defaultEnvFiles, pleaseEnvs...)
	pleaseCmd.Flags().StringArrayVarP(&pleaseValues, "value", "v", []string{}, "values (e.g., '-v value.yaml' or '-v key=val')")
	pleaseValues = append(defaultValueFiles, pleaseValues...)
	pleaseCmd.Flags().StringVarP(&pleaseDecor, "decoration", "d", defaultDecoration, "decoration")
	pleaseInteractive = pleaseCmd.Flags().BoolP("interactive", "i", false, "interactive mode")
	pleaseExplain = pleaseCmd.Flags().BoolP("explain", "x", false, "explain instead of execute")
	pleaseUsePreviousValues = pleaseCmd.Flags().BoolP("previous", "p", false, "load previous values")
	pleaseTerminate = pleaseCmd.Flags().BoolP("terminate", "t", false, "terminate after complete")
	pleaseShowLogTime = pleaseCmd.Flags().BoolP("showLogTime", "s", defaultShowLogTime, "show log time (e.g., '-s false').")
	pleaseCmd.Flags().StringVarP(&pleaseWait, "wait", "w", "0s", "termination waiting duration (e.g., '-w 5s'). Only take effect if -t or --terminate is set")
}

func getDefaultEnvFiles(dir, zarubaEnv string) []string {
	defaultEnvFileCandidates := []string{
		filepath.Join(dir, fmt.Sprintf("%s.env", zarubaEnv)),
	}
	defaultEnvFiles := []string{}
	for _, defaultEnvFileCandidate := range defaultEnvFileCandidates {
		if _, err := os.Stat(defaultEnvFileCandidate); !os.IsNotExist(err) {
			defaultEnvFiles = append(defaultEnvFiles, defaultEnvFileCandidate)
		}
	}
	return defaultEnvFiles
}

func getDefaultValueFiles(dir, zarubaEnv string) []string {
	defaultValueFileCandidates := []string{
		filepath.Join(dir, "default.values.yaml"),
		filepath.Join(dir, fmt.Sprintf("%s.values.yaml", zarubaEnv)),
	}
	defaultValueFiles := []string{}
	for _, defaultValueFileCandidate := range defaultValueFileCandidates {
		if _, err := os.Stat(defaultValueFileCandidate); !os.IsNotExist(err) {
			defaultValueFiles = append(defaultValueFiles, defaultValueFileCandidate)
		}
	}
	return defaultValueFiles
}

func getDefaultProjectFile(dir string) string {
	defaultProjectFile := "${ZARUBA_HOME}/core.zaruba.yaml"
	projectPath, err := helper.GetProjectPath(dir)
	if err != nil {
		return defaultProjectFile
	}
	projectFilePath := filepath.Join(projectPath, "index.zaruba.yaml")
	if _, err := os.Stat(projectFilePath); err == nil {
		return projectFilePath
	}
	projectFilePath = filepath.Join(projectPath, "index.zaruba.yml")
	if _, err := os.Stat(projectFilePath); err == nil {
		return projectFilePath
	}
	return defaultProjectFile
}

func showLastPleaseCommand(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration) {
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
	cmdArgs := cmd.Flags().Args()
	cmdFlags := []string{}
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		key := flag.Name
		val := flag.Value.String()
		cmdFlags = append(cmdFlags, fmt.Sprintf("--%s=%s", key, strUtil.EscapeShellValue(val)))
	})
	logger.Fprintf(os.Stderr, "%s%s %s %s %s\n", decoration.Yellow, commandName, strings.Join(cmdArgs, " "), strings.Join(cmdFlags, " "), decoration.Normal)
}

func getTaskNameInteractivelyOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter) (taskName string) {
	taskName, err := prompter.GetTaskName()
	if err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
	return taskName
}

func getActionOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskName string) (action *input.Action) {
	action, err := prompter.GetAction(taskName)
	if err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
	return action
}

func explainOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, explainer *explainer.Explainer, taskNames []string) {
	if err := explainer.Explain(taskNames...); err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
}

func loadPreviousValuesOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, project *dsl.Project, previousValueFile string) {
	if err := previousval.Load(project, previousValueFile); err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
}

func askProjectValuesByTasksOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskNames []string) {
	if err := prompter.SetProjectValuesByTask(taskNames); err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
}

func askProjectEnvOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskNames []string) {
	if err := prompter.GetAdditionalEnv(taskNames); err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
}

func askProjectValueOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter) {
	if err := prompter.GetAdditionalValue(); err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
}

func askAutoTerminateOrExit(cmd *cobra.Command, logger *output.ConsoleLogger, decoration *output.Decoration, prompter *input.Prompter, taskNames []string) (autoTerminate bool) {
	autoTerminate, err := prompter.GetAutoTerminate(taskNames)
	if err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
	return autoTerminate
}

func initProjectOrExit(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, project *dsl.Project) {
	if err := project.Init(); err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
}

func getProjectAndTaskName(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, showLogTime bool, args []string) (project *dsl.Project, taskNames []string) {
	project, err := dsl.NewProject(pleaseProjectFile, decoration, showLogTime)
	if err != nil {
		cmdHelper.Exit(cmd, logger, decoration, err)
	}
	taskNames = []string{}
	//  distinguish taskNames and additional values
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			if err = project.AddValue(arg); err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
			continue
		}
		taskNames = append(taskNames, arg)
	}
	// process envs
	for _, env := range pleaseEnvs {
		if err = project.AddEnv(env); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	}
	// process values from flag
	for _, value := range pleaseValues {
		if err = project.AddValue(value); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	}
	return project, taskNames
}
