package explainer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/logger"
)

// Explain explain inputs or tasks exclusively if the first element of the keywords is either "input" or "task". Otherwise, both inputs and tasks will be explained
func Explain(project *config.Project, keywords []string) (err error) {
	// first element of "keywords" is "input", only show explain inputs
	if len(keywords) >= 1 && keywords[0] == "input" {
		keyword := strings.Join(keywords[1:], " ")
		return ExplainInputs(project, keyword)
	}
	// first element of "keywords" is "task", only show explain tasks
	if len(keywords) >= 1 && keywords[0] == "task" {
		keyword := strings.Join(keywords[1:], " ")
		return ExplainTasks(project, keyword)
	}
	// explain tasks and inputs
	keyword := strings.Join(keywords, " ")
	if err = ExplainTasks(project, keyword); err != nil {
		return err
	}
	return ExplainInputs(project, keyword)
}

// ExplainInputs explain all inputs
func ExplainInputs(project *config.Project, keyword string) (err error) {
	if !project.IsInitialized {
		return fmt.Errorf("Cannot explain inputs because project was not initialize")
	}
	r := getRegexKeyword(keyword)
	d := logger.NewDecoration()
	inputIndentation := strings.Repeat(" ", 6)
	inputFieldIndentation := inputIndentation + strings.Repeat(" ", 2)
	totalMatch := 0
	for _, inputName := range project.GetSortedInputNames() {
		if !r.MatchString(inputName) {
			continue
		}
		if totalMatch == 0 {
			logger.Printf("%sINPUT VARIABLES:%s\n", d.Yellow, d.Normal)
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
	logger.Printf("%d input variable(s) matched '%s' keyword.\n", totalMatch, keyword)
	logger.Printf("You can also use %s%szaruba please explain input <keyword>%s to refine the result.\n", d.Bold, d.Yellow, d.Normal)
	return nil
}

// ExplainTasks explain all tasks matching the keyword
func ExplainTasks(project *config.Project, keyword string) (err error) {
	if !project.IsInitialized {
		return fmt.Errorf("Cannot explain tasks because project was not initialize")
	}
	r := getRegexKeyword(keyword)
	d := logger.NewDecoration()
	unpublishedMatch := showTasks(project, false, r)
	publishedMatch := showTasks(project, true, r)
	totalMatch := unpublishedMatch + publishedMatch
	logger.Printf("%d task(s) matched '%s' keyword:\n", totalMatch, keyword)
	logger.Printf("    %d base task(s)\n", unpublishedMatch)
	logger.Printf("    %d published task(s)\n", publishedMatch)
	logger.Printf("You can also use %s%szaruba please explain task <keyword>%s to refine the result.\n", d.Bold, d.Yellow, d.Normal)
	return nil
}

func showTasks(project *config.Project, published bool, r *regexp.Regexp) (totalMatch int) {
	d := logger.NewDecoration()
	taskIndentation := strings.Repeat(" ", 6)
	taskFieldIndentation := taskIndentation + strings.Repeat(" ", 5)
	taskPrefix := ""
	if published {
		taskPrefix = "zaruba please "
	}
	for _, taskName := range project.GetSortedTaskNames() {
		task := project.Tasks[taskName]
		if (task.Private && published) || (!task.Private && !published) {
			continue
		}
		if !r.MatchString(taskName) {
			continue
		}
		if totalMatch == 0 {
			if published {
				logger.Printf("%sPUBLISHED TASKS:%s\n", d.Yellow, d.Normal)
			} else {
				logger.Printf("%sBASE TASKS:%s\n", d.Yellow, d.Normal)
			}
		}
		fmt.Printf("%s%s %s%s%s%s%s\n", taskIndentation, task.Icon, d.Yellow, d.Bold, taskPrefix, task.GetName(), d.Normal)
		showField(taskFieldIndentation, "PUBLISHED", fmt.Sprintf("%t", !task.Private), true)
		showField(taskFieldIndentation, "DECLARED ON", task.GetFileLocation(), true)
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
	inputs, inputOrder, err := task.Project.GetInputs([]string{task.GetName()})
	if err == nil && len(inputOrder) > 0 {
		showField(fieldIndentation, "PARAMETERS", "", true)
		d := logger.NewDecoration()
		for _, inputName := range inputOrder {
			input := inputs[inputName]
			inputCaption := fmt.Sprintf("- %s", inputName)
			if input.DefaultValue != "" {
				inputCaption = fmt.Sprintf("%s %s(default:%s %s%s%s)%s", inputCaption, d.Faint, d.Normal, d.Yellow, input.DefaultValue, d.Faint, d.Normal)
			}
			fmt.Printf("%s%s\n", fieldIndentation, inputCaption)
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
