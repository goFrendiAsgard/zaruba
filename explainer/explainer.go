package explainer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/monitor"
)

// Explain explain inputs or tasks exclusively if the first element of the keywords is either "input" or "task". Otherwise, both inputs and tasks will be explained
func Explain(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project, keywords []string) (err error) {
	// first element of "keywords" is "input", only show explain inputs
	if len(keywords) >= 1 && keywords[0] == "input" {
		keyword := strings.Join(keywords[1:], " ")
		return explainInputs(logger, decoration, project, keyword)
	}
	// first element of "keywords" is "task", only show explain tasks
	if len(keywords) >= 1 && keywords[0] == "task" {
		keyword := strings.Join(keywords[1:], " ")
		return explainTasks(logger, decoration, project, keyword)
	}
	// explain tasks and inputs
	keyword := strings.Join(keywords, " ")
	if err = explainTasks(logger, decoration, project, keyword); err != nil {
		return err
	}
	return explainInputs(logger, decoration, project, keyword)
}

// explainInputs explain all inputs
func explainInputs(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project, keyword string) (err error) {
	if !project.IsInitialized {
		return fmt.Errorf("Cannot explain inputs because project was not initialize")
	}
	r := getRegexKeyword(logger, decoration, keyword)
	inputIndentation := strings.Repeat(" ", 6)
	inputFieldIndentation := inputIndentation + strings.Repeat(" ", 2)
	totalMatch := 0
	for _, inputName := range project.GetSortedInputNames() {
		if !r.MatchString(inputName) {
			continue
		}
		if totalMatch == 0 {
			logger.DPrintf("%sINPUT VARIABLES:%s\n", decoration.Yellow, decoration.Normal)
		}
		input := project.Inputs[inputName]
		logger.Printf("%s%s%s%s%s\n", inputIndentation, decoration.Yellow, decoration.Bold, inputName, decoration.Normal)
		if input.DefaultValue != "" {
			showField(logger, decoration, inputFieldIndentation, "DEFAULT", input.DefaultValue, false)
		} else {
			showField(logger, decoration, inputFieldIndentation, "DEFAULT", "empty", true)
		}
		if input.Description != "" {
			showField(logger, decoration, inputFieldIndentation, "DESCRIPTION", input.Description, false)
		}
		totalMatch++
	}
	logger.DPrintf("%d input variable(s) matched '%s' keyword.\n", totalMatch, keyword)
	logger.DPrintf("You can also use %s%szaruba please explain input <keyword>%s to refine the result.\n", decoration.Bold, decoration.Yellow, decoration.Normal)
	return nil
}

// explainTasks explain all tasks matching the keyword
func explainTasks(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project, keyword string) (err error) {
	if !project.IsInitialized {
		return fmt.Errorf("Cannot explain tasks because project was not initialize")
	}
	r := getRegexKeyword(logger, decoration, keyword)
	unpublishedMatch := showTasks(logger, decoration, project, false, r)
	publishedMatch := showTasks(logger, decoration, project, true, r)
	totalMatch := unpublishedMatch + publishedMatch
	logger.DPrintf("%d task(s) matched '%s' keyword:\n", totalMatch, keyword)
	logger.DPrintf("    %d base task(s)\n", unpublishedMatch)
	logger.DPrintf("    %d published task(s)\n", publishedMatch)
	logger.DPrintf("You can also use %s%szaruba please explain task <keyword>%s to refine the result.\n", decoration.Bold, decoration.Yellow, decoration.Normal)
	return nil
}

func showTasks(logger monitor.Logger, decoration *monitor.Decoration, project *config.Project, published bool, r *regexp.Regexp) (totalMatch int) {
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
				logger.DPrintf("%sPUBLISHED TASKS:%s\n", decoration.Yellow, decoration.Normal)
			} else {
				logger.DPrintf("%sBASE TASKS:%s\n", decoration.Yellow, decoration.Normal)
			}
		}
		logger.Printf("%s%s %s%s%s%s%s\n", taskIndentation, task.Icon, decoration.Yellow, decoration.Bold, taskPrefix, task.GetName(), decoration.Normal)
		showField(logger, decoration, taskFieldIndentation, "PUBLISHED", fmt.Sprintf("%t", !task.Private), true)
		showField(logger, decoration, taskFieldIndentation, "DECLARED ON", task.GetFileLocation(), true)
		showTaskParameters(logger, decoration, task, taskFieldIndentation)
		showTaskDescription(logger, decoration, task, taskFieldIndentation)
		showTaskDependencies(logger, decoration, task, taskFieldIndentation)
		showTaskExtend(logger, decoration, task, taskFieldIndentation)
		totalMatch++
	}
	return totalMatch
}

func showField(logger monitor.Logger, decoration *monitor.Decoration, fieldIndentation string, fieldName string, value string, faint bool) {
	paddedFieldName := fmt.Sprintf("%-15v", fieldName)
	trimmedValue := strings.Trim(value, "\n ")
	rows := strings.Split(trimmedValue, "\n")
	if len(rows) == 1 {
		decoratedValue := trimmedValue
		if faint {
			decoratedValue = fmt.Sprintf("%s%s%s", decoration.Faint, trimmedValue, decoration.Normal)
		}
		logger.Printf("%s%s%s:%s %s\n", fieldIndentation, decoration.Blue, paddedFieldName, decoration.Normal, decoratedValue)
		return
	}
	logger.Printf("%s%s%s:%s\n", fieldIndentation, decoration.Blue, paddedFieldName, decoration.Normal)
	for _, row := range rows {
		if faint {
			logger.Printf("%s  %s%s%s\n", fieldIndentation, decoration.Faint, row, decoration.Normal)
			continue
		}
		logger.Printf("%s  %s\n", fieldIndentation, row)
	}
}

func showTaskDependencies(logger monitor.Logger, decoration *monitor.Decoration, task *config.Task, fieldIndentation string) {
	if len(task.Dependencies) > 0 {
		showField(logger, decoration, fieldIndentation, "DEPENDENCIES", strings.Join(task.Dependencies, ", "), true)
	}
}

func showTaskExtend(logger monitor.Logger, decoration *monitor.Decoration, task *config.Task, fieldIndentation string) {
	if task.Extend != "" {
		showField(logger, decoration, fieldIndentation, "EXTENDED FROM", task.Extend, true)
	}
}

func showTaskParameters(logger monitor.Logger, decoration *monitor.Decoration, task *config.Task, fieldIndentation string) {
	inputs, inputOrder, err := task.Project.GetInputs([]string{task.GetName()})
	if err == nil && len(inputOrder) > 0 {
		showField(logger, decoration, fieldIndentation, "PARAMETERS", "", true)
		for _, inputName := range inputOrder {
			input := inputs[inputName]
			inputCaption := fmt.Sprintf("- %s", inputName)
			if input.DefaultValue != "" {
				inputCaption = fmt.Sprintf("%s %s(default:%s %s%s%s)%s", inputCaption, decoration.Faint, decoration.Normal, decoration.Yellow, input.DefaultValue, decoration.Faint, decoration.Normal)
			}
			logger.Printf("%s%s\n", fieldIndentation, inputCaption)
			if input.Description != "" {
				description := strings.Trim(input.Description, "\n ")
				rows := strings.Split(description, "\n")
				for _, row := range rows {
					decoratedRow := fmt.Sprintf("%s%s%s", decoration.Faint, row, decoration.Normal)
					logger.Printf("%s    %s\n", fieldIndentation, decoratedRow)
				}
			}
		}
	}
}

func showTaskDescription(logger monitor.Logger, decoration *monitor.Decoration, task *config.Task, fieldIndentation string) {
	if task.Description != "" {
		showField(logger, decoration, fieldIndentation, "DESCRIPTION", task.Description, false)
	}
}

func getRegexKeyword(logger monitor.Logger, decoration *monitor.Decoration, searchPattern string) (r *regexp.Regexp) {
	if searchPattern == "" {
		r, _ = regexp.Compile(".*")
		return r
	}
	r, err := regexp.Compile("(?i)" + searchPattern)
	if err != nil {
		logger.DPrintfError("Invalid regex: %s\n", searchPattern)
		r, _ = regexp.Compile(".*")
	}
	return r
}
