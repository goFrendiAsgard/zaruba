package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/logger"
	"github.com/state-alchemists/zaruba/runner"
)

var pleaseEnv []string
var pleaseKwargs []string
var pleaseFile string

// pleaseCmd represents the please command
var pleaseCmd = &cobra.Command{
	Use:   "please",
	Short: "Ask Zaruba to do something for you",
	Long:  "💀 Ask Zaruba to do something for you",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.NewConfig(pleaseFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		// process globalEnv
		for _, env := range pleaseEnv {
			conf.AddGlobalEnv(env)
		}
		// process kwargs from flag
		for _, kwarg := range pleaseKwargs {
			if err = conf.AddKwargs(kwarg); err != nil {
				fmt.Println(err)
				return
			}
		}
		//  distinguish taskNames and additional kwargs
		taskNames := []string{}
		for _, arg := range args {
			if strings.Contains(arg, "=") {
				conf.AddKwargs(arg)
				continue
			}
			taskNames = append(taskNames, arg)
		}
		// init
		if err = conf.Init(); err != nil {
			fmt.Println(err)
			return
		}
		// handle special cases
		if handleSpecialCases(conf, taskNames) {
			return
		}
		// run
		r := runner.NewRunner(conf, taskNames, time.Minute*5)
		if err := r.Run(); err != nil {
			fmt.Println(err)
		}
	},
}

func handleSpecialCases(conf *config.ProjectConfig, taskNames []string) (handled bool) {
	d := logger.NewDecoration()
	// special case: task is not given
	if len(taskNames) == 0 {
		logger.Printf("%sPlease what?%s\n", d.Bold, d.Normal)
		logger.Printf("Here are some possible tasks you can execute:\n")
		showTasks(conf, true)
		return true
	}
	if len(taskNames) > 1 {
		return false
	}
	taskName := taskNames[0]
	_, taskIsDeclaredByUser := conf.Tasks[taskName]
	if taskIsDeclaredByUser {
		return false
	}
	specialTasks := map[string](func()){
		"showTasks": func() {
			fmt.Printf("\n%s%s%s\n", d.Bold, "PUBLISHED TASKS", d.Normal)
			showTasks(conf, true)
			fmt.Printf("\n%s%s%s\n", d.Bold, "UNPUBLISHED TASKS", d.Normal)
			showTasks(conf, false)
		},
		"showPublishedTasks": func() {
			showTasks(conf, true)
		},
		"showUnpublishedTasks": func() {
			showTasks(conf, false)
		},
	}
	action, taskIsSpecial := specialTasks[taskName]
	if !taskIsSpecial {
		return false
	}
	// special case: task is given, it is special task, and no other tasks exist
	action()
	return true
}

func showTasks(conf *config.ProjectConfig, showPublished bool) {
	d := logger.NewDecoration()
	taskIndentation := strings.Repeat(" ", 6)
	taskFieldIndentation := taskIndentation + strings.Repeat(" ", 5)
	taskPrefix := ""
	if showPublished {
		taskPrefix = "zaruba please "
	}
	publishedTask := conf.GetPublishedTask()
	for _, taskName := range conf.SortedTaskNames {
		if _, exist := publishedTask[taskName]; (exist && showPublished) || (!exist && !showPublished) {
			task := conf.Tasks[taskName]
			fmt.Printf("%s%s %s%s%s%s%s\n", taskIndentation, task.Icon, d.Yellow, taskPrefix, d.Bold, task.Name, d.Normal)
			fmt.Printf("%s%sDECLARED ON:%s%s %s%s\n", taskFieldIndentation, d.Blue, d.Normal, d.Faint, task.FileLocation, d.Normal)
			showTaskDescription(task, taskFieldIndentation)
		}
	}
}

func showTaskDescription(task *config.Task, fieldIndentation string) {
	if task.Description != "" {
		d := logger.NewDecoration()
		description := strings.TrimSpace(task.Description)
		rows := strings.Split(description, "\n")
		for index, row := range rows {
			if index == 0 {
				row = fmt.Sprintf("%sDESCRIPTION:%s %s", d.Blue, d.Normal, row)
			}
			fmt.Printf("%s%s\n", fieldIndentation, row)
		}
	}
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
	// define defaultPleaseKwargs
	defaultPleaseKwargs := []string{}
	defaultKwargsFile := filepath.Join(dir, "default.kwargs.yaml")
	if _, err := os.Stat(defaultKwargsFile); !os.IsNotExist(err) {
		defaultPleaseKwargs = append(defaultPleaseKwargs, defaultKwargsFile)
	}
	// define defaultEnvFile
	defaultEnv := []string{}
	defaultEnvFile := filepath.Join(dir, ".env")
	if _, err := os.Stat(defaultEnvFile); !os.IsNotExist(err) {
		defaultEnv = append(defaultEnv, defaultEnvFile)
	}
	// register flags
	pleaseCmd.Flags().StringVarP(&pleaseFile, "file", "f", defaultPleaseFile, "custom file")
	pleaseCmd.Flags().StringArrayVarP(&pleaseEnv, "environment", "e", defaultEnv, "environment file or pairs (e.g: '-e environment.env' or '-e key=val')")
	pleaseCmd.Flags().StringArrayVarP(&pleaseKwargs, "kwargs", "k", defaultPleaseKwargs, "yaml file or pairs (e.g: '-k value.yaml' or '-k key=val')")
}
