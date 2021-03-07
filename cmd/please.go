package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/explainer"
	"github.com/state-alchemists/zaruba/logger"
	"github.com/state-alchemists/zaruba/runner"
)

var pleaseEnv []string
var pleaseValues []string
var pleaseFile string
var pleaseInteractive *bool

// pleaseCmd represents the please command
var pleaseCmd = &cobra.Command{
	Use:     "please",
	Short:   "Ask Zaruba to do something for you",
	Long:    "💀 Ask Zaruba to do something for you",
	Aliases: []string{"run", "do", "invoke", "perform"},
	Run: func(cmd *cobra.Command, args []string) {
		project, taskNames, err := getProjectAndTaskNames(args)
		if err != nil {
			showErrorAndExit(err)
		}
		// init
		if err = project.Init(); err != nil {
			showErrorAndExit(err)
		}
		d := logger.NewDecoration()
		// no task provided
		if len(taskNames) == 0 {
			logger.Printf("%sPlease what?%s\n", d.Bold, d.Normal)
			logger.Printf("Here are several things you can try:\n")
			logger.Printf("* %szaruba please explain task %s%s[task-keyword]%s\n", d.Yellow, d.Normal, d.Blue, d.Normal)
			logger.Printf("* %szaruba please explain input %s%s[input-keyword]%s\n", d.Yellow, d.Normal, d.Blue, d.Normal)
			logger.Printf("* %szaruba please explain %s%s[task-or-input-keyword]%s\n", d.Yellow, d.Normal, d.Blue, d.Normal)
			return
		}
		// handle "please explain"
		if taskNames[0] == "explain" {
			if len(taskNames) >= 2 {
				if taskNames[1] == "input" || taskNames[1] == "task" {
					keyword := strings.Join(taskNames[2:], " ")
					// handle "please explain input"
					if taskNames[1] == "input" {
						explainer.ExplainInputs(project, keyword)
						return
					}
					// handle "please explain task"
					explainer.ExplainTasks(project, keyword)
					return
				}
			}
			// handle "please explain"
			keyword := strings.Join(taskNames[1:], " ")
			explainer.ExplainTasks(project, keyword)
			explainer.ExplainInputs(project, keyword)
			return
		}
		if *pleaseInteractive {
			askInputs(project, taskNames)
		}
		// run
		r := runner.NewRunner(project, taskNames, time.Minute*5)
		if err := r.Run(); err != nil {
			showErrorAndExit(err)
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
	pleaseCmd.Flags().StringVarP(&pleaseFile, "file", "f", defaultPleaseFile, "task file")
	pleaseCmd.Flags().StringArrayVarP(&pleaseEnv, "environment", "e", defaultEnv, "environment file or pairs (e.g: '-e environment.env' or '-e key=val')")
	pleaseCmd.Flags().StringArrayVarP(&pleaseValues, "value", "v", defaultPleaseValues, "yaml file or pairs (e.g: '-v value.yaml' or '-v key=val')")
	pleaseInteractive = pleaseCmd.Flags().BoolP("interactive", "i", false, "if set, zaruba will ask you to fill inputs (e.g: -i)")
}

func askInputs(project *config.Project, taskNames []string) (err error) {
	inputs, inputOrder, err := project.GetInputs(taskNames)
	d := logger.NewDecoration()
	inputCount := len(inputOrder)
	if inputCount == 0 {
		logger.Printf("%sZaruba is running in interactive mode. But no value you can set interactively.%s\n", d.Yellow, d.Normal)
		return err
	}
	logger.Printf("%sZaruba is running in interactive mode. You will be able to set some values interactively.%s\n", d.Yellow, d.Normal)
	logger.Printf("%sLeave blank if you want to keep current values.%s\n", d.Yellow, d.Normal)
	for inputIndex, inputName := range inputOrder {
		input := inputs[inputName]
		if input.Description != "" {
			fmt.Printf("\n%s%s%s%s\n", d.Bold, d.Blue, inputName, d.Normal)
			fmt.Printf("%s\n\n", strings.Trim(input.Description, "\n "))
		}
		decoratedValue := "empty"
		value := project.Values[inputName]
		if value != "" {
			decoratedValue = fmt.Sprintf("%s%s%s", d.Yellow, value, d.Normal)
		}
		decoratedIndex := fmt.Sprintf("%s%d of %d)%s", d.Faint, inputIndex+1, inputCount, d.Normal)
		fmt.Printf("%s %s (currently %s): ", decoratedIndex, inputName, decoratedValue)
		userValue := ""
		fmt.Scanf("%s", &userValue)
		if userValue != "" {
			project.SetValue(inputName, userValue)
		}
	}
	return err
}

func getProjectAndTaskNames(args []string) (project *config.Project, taskNames []string, err error) {
	taskNames = []string{}
	project, err = config.NewProject(pleaseFile)
	if err != nil {
		fmt.Println(err)
		return project, taskNames, err
	}
	// process globalEnv
	for _, env := range pleaseEnv {
		project.AddGlobalEnv(env)
	}
	// process values from flag
	for _, value := range pleaseValues {
		if err = project.AddValues(value); err != nil {
			fmt.Println(err)
			return project, taskNames, err
		}
	}
	//  distinguish taskNames and additional values
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			project.AddValues(arg)
			continue
		}
		_, argIsTask := project.Tasks[arg]
		if !argIsTask {
			if arg == "autostop" {
				project.AddValues("autostop=true")
				continue
			}
		}
		taskNames = append(taskNames, arg)
	}
	return project, taskNames, err
}

func showErrorAndExit(err error) {
	d := logger.NewDecoration()
	if err != nil {
		logger.PrintfError("%s%s%s%s\n", d.Bold, d.Red, err.Error(), d.Normal)
		os.Exit(1)
	}
}
