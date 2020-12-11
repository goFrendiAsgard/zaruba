package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
	Use:     "please",
	Short:   "Ask Zaruba to do something for you",
	Long:    "💀 Ask Zaruba to do something for you",
	Aliases: []string{"run", "do", "invoke", "perform"},
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
			_, argIsTask := conf.Tasks[arg]
			if !argIsTask {
				if arg == "autostop" {
					conf.AddKwargs("onComplete=stop")
					continue
				}
			}
			taskNames = append(taskNames, arg)
		}
		// init
		d := logger.NewDecoration()
		if err = conf.Init(); err != nil {
			logger.PrintfError("%s%s%s%s\n", d.Bold, d.Red, err.Error(), d.Normal)
			os.Exit(1)
		}
		// handle special cases
		if handleShowTask(conf, taskNames) {
			return
		}
		// run
		r := runner.NewRunner(conf, taskNames, time.Minute*5)
		if err := r.Run(); err != nil {
			logger.PrintfError("%s%s%s%s\n", d.Bold, d.Red, err.Error(), d.Normal)
			os.Exit(1)
		}
	},
}

func handleShowTask(conf *config.ProjectConfig, taskNames []string) (handled bool) {
	d := logger.NewDecoration()
	// special case: task is not given
	if len(taskNames) == 0 {
		logger.Printf("%sPlease what?%s\n", d.Bold, d.Normal)
		logger.Printf("Here are some possible tasks you can execute:\n")
		totalMatched := showTasks(conf, true, getRegexSearchPattern(""))
		showSearchFooter(totalMatched, "")
		return true
	}
	taskName, keyword := taskNames[0], strings.Join(taskNames[1:], " ")
	_, taskDeclared := conf.Tasks[taskName]
	if taskDeclared || taskName != "explain" {
		return false
	}
	r := getRegexSearchPattern(keyword)
	published, publishExist := conf.Kwargs["published"]
	totalMatched := 0
	if !publishExist || published == "false" {
		totalMatched += showTasks(conf, false, r)
	}
	if !publishExist || published == "true" {
		totalMatched += showTasks(conf, true, r)
	}
	showSearchFooter(totalMatched, keyword)
	return true
}

func showSearchFooter(totalMatched int, keyword string) {
	d := logger.NewDecoration()
	logger.Printf("%d task(s) matched '%s' keyword.\n", totalMatched, keyword)
	logger.Printf("You can also use %s%szaruba please explain <keyword>%s to refine the result.\n", d.Bold, d.Yellow, d.Normal)
}

func getRegexSearchPattern(searchPattern string) (r *regexp.Regexp) {
	if searchPattern == "" {
		r, _ = regexp.Compile(".*")
		return r
	}
	r, err := regexp.Compile("(?i)" + searchPattern)
	if err != nil {
		logger.PrintfError("Invalid regex: %s\n", searchPattern)
		r, _ = regexp.Compile(".*")
	}
	return r
}

func showTasks(conf *config.ProjectConfig, showPublished bool, r *regexp.Regexp) (totalMatch int) {
	d := logger.NewDecoration()
	taskIndentation := strings.Repeat(" ", 6)
	taskFieldIndentation := taskIndentation + strings.Repeat(" ", 5)
	taskPrefix := ""
	if showPublished {
		taskPrefix = "zaruba please "
	}
	for _, taskName := range conf.SortedTaskNames {
		task := conf.Tasks[taskName]
		if (task.Private && showPublished) || (!task.Private && !showPublished) {
			continue
		}
		if !r.MatchString(taskName) {
			continue
		}
		fmt.Printf("%s%s %s%s%s%s%s\n", taskIndentation, task.Icon, d.Yellow, taskPrefix, d.Bold, task.Name, d.Normal)
		fmt.Printf("%s%sPUBLISHED  :%s%s %t%s\n", taskFieldIndentation, d.Blue, d.Normal, d.Faint, !task.Private, d.Normal)
		fmt.Printf("%s%sDECLARED ON:%s%s %s%s\n", taskFieldIndentation, d.Blue, d.Normal, d.Faint, task.FileLocation, d.Normal)
		showTaskDescription(task, taskFieldIndentation)
		totalMatch++
	}
	return totalMatch
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
