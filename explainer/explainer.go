package explainer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/logger"
)

// ExplainInputs explain all inputs
func ExplainInputs(project *config.Project, keyword string) {
	r := getRegexKeyword(keyword)
	d := logger.NewDecoration()
	inputIndentation := strings.Repeat(" ", 6)
	inputFieldIndentation := inputIndentation + strings.Repeat(" ", 2)
	totalMatch := 0
	for _, inputName := range project.SortedInputNames {
		if !r.MatchString(inputName) {
			continue
		}
		input := project.Inputs[inputName]
		fmt.Printf("%s%s%s%s%s\n", inputIndentation, d.Yellow, d.Bold, inputName, d.Normal)
		if input.DefaultValue != "" {
			showField(inputFieldIndentation, "DEFAULT", input.DefaultValue, false)
		} else {
			showField(inputFieldIndentation, "DEFAULT", "empty", true)
		}
		if input.Description != "" {
			showField(inputFieldIndentation, "DESCRIPTION", input.Description, false)
		}
		totalMatch++
	}
	logger.Printf("%d input(s) matched '%s' keyword.\n", totalMatch, keyword)
	logger.Printf("You can also use %s%szaruba please explain inputs <keyword>%s to refine the result.\n", d.Bold, d.Yellow, d.Normal)
}

// ExplainTasks explain all tasks matching the keyword
func ExplainTasks(project *config.Project, keyword string) {
	r := getRegexKeyword(keyword)
	d := logger.NewDecoration()
	totalMatch := showTasks(project, false, r)
	totalMatch += showTasks(project, true, r)
	logger.Printf("%d task(s) matched '%s' keyword.\n", totalMatch, keyword)
	logger.Printf("You can also use %s%szaruba please explain <keyword>%s to refine the result.\n", d.Bold, d.Yellow, d.Normal)
}

func showTasks(project *config.Project, published bool, r *regexp.Regexp) (totalMatch int) {
	d := logger.NewDecoration()
	taskIndentation := strings.Repeat(" ", 6)
	taskFieldIndentation := taskIndentation + strings.Repeat(" ", 5)
	taskPrefix := ""
	if published {
		taskPrefix = "zaruba please "
	}
	for _, taskName := range project.SortedTaskNames {
		task := project.Tasks[taskName]
		if (task.Private && published) || (!task.Private && !published) {
			continue
		}
		if !r.MatchString(taskName) {
			continue
		}
		fmt.Printf("%s%s %s%s%s%s%s\n", taskIndentation, task.Icon, d.Yellow, d.Bold, taskPrefix, task.Name, d.Normal)
		showField(taskFieldIndentation, "PUBLISHED", fmt.Sprintf("%t", !task.Private), true)
		showField(taskFieldIndentation, "DECLARED ON", task.FileLocation, true)
		showTaskParameters(task, taskFieldIndentation)
		showTaskDescription(task, taskFieldIndentation)
		showTaskDependencies(task, taskFieldIndentation)
		showTaskExtend(task, taskFieldIndentation)
		totalMatch++
	}
	return totalMatch
}

func showField(fieldIndentation string, fieldName string, value string, faint bool) {
	d := logger.NewDecoration()
	paddedFieldName := fmt.Sprintf("%-15v", fieldName)
	trimmedValue := strings.Trim(value, "\n ")
	rows := strings.Split(trimmedValue, "\n")
	if len(rows) == 1 {
		decoratedValue := trimmedValue
		if faint {
			decoratedValue = fmt.Sprintf("%s%s%s", d.Faint, trimmedValue, d.Normal)
		}
		fmt.Printf("%s%s%s:%s %s\n", fieldIndentation, d.Blue, paddedFieldName, d.Normal, decoratedValue)
		return
	}
	fmt.Printf("%s%s%s:%s\n", fieldIndentation, d.Blue, paddedFieldName, d.Normal)
	for _, row := range rows {
		if faint {
			fmt.Printf("%s  %s%s%s\n", fieldIndentation, d.Faint, row, d.Normal)
			continue
		}
		fmt.Printf("%s  %s\n", fieldIndentation, row)
	}
}

func showTaskDependencies(task *config.Task, fieldIndentation string) {
	if len(task.Dependencies) > 0 {
		showField(fieldIndentation, "DEPENDENCIES", strings.Join(task.Dependencies, ", "), true)
	}
}

func showTaskExtend(task *config.Task, fieldIndentation string) {
	if task.Extend != "" {
		showField(fieldIndentation, "EXTENDED FROM", task.Extend, true)
	}
}

func showTaskParameters(task *config.Task, fieldIndentation string) {
	if len(task.Inputs) > 0 {
		showField(fieldIndentation, "PARAMETERS", "", true)
		d := logger.NewDecoration()
		for _, inputName := range task.Inputs {
			input := task.Project.Inputs[inputName]
			inputCaption := fmt.Sprintf("- %s", inputName)
			if input.DefaultValue != "" {
				inputCaption = fmt.Sprintf("%s %s(default:%s %s%s%s)%s", inputCaption, d.Faint, d.Normal, d.Yellow, input.DefaultValue, d.Faint, d.Normal)
			}
			fmt.Printf("%s  %s\n", fieldIndentation, inputCaption)
			if input.Description != "" {
				description := strings.Trim(input.Description, "\n ")
				rows := strings.Split(description, "\n")
				for _, row := range rows {
					decoratedRow := fmt.Sprintf("%s%s%s", d.Faint, row, d.Normal)
					fmt.Printf("%s    %s\n", fieldIndentation, decoratedRow)
				}
			}
		}
	}
}

func showTaskDescription(task *config.Task, fieldIndentation string) {
	if task.Description != "" {
		showField(fieldIndentation, "DESCRIPTION", task.Description, false)
	}
}

func getRegexKeyword(searchPattern string) (r *regexp.Regexp) {
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
