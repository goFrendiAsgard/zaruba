package explainer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/logger"
)

// ExplainTasks explain all private tasks matching the keyword
func ExplainTasks(project *config.Project, keyword string) {
	r := getRegexKeyword(keyword)
	totalMatch := showTasks(project, false, r)
	totalMatch += showTasks(project, true, r)
	showFooter(totalMatch, keyword)
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
		showTaskField(taskFieldIndentation, "PUBLISHED", fmt.Sprintf("%t", !task.Private))
		showTaskField(taskFieldIndentation, "DECLARED ON", task.FileLocation)
		showTaskDescription(task, taskFieldIndentation)
		showTaskDependencies(task, taskFieldIndentation)
		showTaskExtend(task, taskFieldIndentation)
		totalMatch++
	}
	return totalMatch
}

func showTaskField(taskFieldIndentation string, fieldName string, value string) {
	d := logger.NewDecoration()
	paddedFieldName := fmt.Sprintf("%-15v", fieldName)
	fmt.Printf("%s%s%s:%s%s %s%s\n", taskFieldIndentation, d.Blue, paddedFieldName, d.Normal, d.Faint, value, d.Normal)
}

func showTaskDependencies(task *config.Task, fieldIndentation string) {
	if len(task.Dependencies) > 0 {
		showTaskField(fieldIndentation, "DEPENDENCIES", strings.Join(task.Dependencies, ", "))
	}
}

func showTaskExtend(task *config.Task, fieldIndentation string) {
	if task.Extend != "" {
		showTaskField(fieldIndentation, "EXTENDED FROM", task.Extend)
	}
}

func showTaskDescription(task *config.Task, fieldIndentation string) {
	if task.Description != "" {
		description := strings.TrimSpace(task.Description)
		rows := strings.Split(description, "\n")
		d := logger.NewDecoration()
		for index, row := range rows {
			if index == 0 {
				showTaskField(fieldIndentation, "DESCRIPTION", row)
				continue
			}
			fmt.Printf("%s  %s%s%s\n", fieldIndentation, d.Faint, row, d.Normal)
		}
	}
}

func showFooter(totalMatched int, keyword string) {
	d := logger.NewDecoration()
	logger.Printf("%d task(s) matched '%s' keyword.\n", totalMatched, keyword)
	logger.Printf("You can also use %s%szaruba please explain <keyword>%s to refine the result.\n", d.Bold, d.Yellow, d.Normal)
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
