package explainer

import (
	"fmt"
	"sort"
	"strings"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
)

type Explainer struct {
	logger  output.Logger
	d       *output.Decoration
	project *config.Project
}

func NewExplainer(logger output.Logger, decoration *output.Decoration, project *config.Project) *Explainer {
	return &Explainer{
		logger:  logger,
		d:       decoration,
		project: project,
	}
}

func (e *Explainer) listToMultiLineStr(list []string) string {
	if len(list) == 0 {
		return ""
	}
	lines := []string{}
	for _, line := range list {
		line = str.Indent(line, "  ")
		lines = append(lines, fmt.Sprintf("- %s", line))
	}
	return strings.Join(lines, "\n")
}

func (e *Explainer) listToStr(list []string) string {
	if len(list) == 0 {
		return ""
	}
	separator := fmt.Sprintf("%s,%s ", e.d.Blue, e.d.Normal)
	return fmt.Sprintf("%s[ %s%s%s ]%s", e.d.Blue, e.d.Normal, strings.Join(list, separator), e.d.Blue, e.d.Normal)
}

func (e *Explainer) getStrOrBlank(str string) string {
	if str == "" {
		return fmt.Sprintf("%sBlank%s", e.d.Blue, e.d.Normal)
	}
	return str
}

func (e *Explainer) getFieldKeys(list []string) (keys []string) {
	keys = []string{}
	maxLength := 0
	for _, key := range list {
		if len(key) > maxLength {
			maxLength = len(key)
		}
	}
	for _, key := range list {
		fieldKey := key + strings.Repeat(" ", maxLength-len(key))
		keys = append(keys, fieldKey)
	}
	return keys
}

func (e *Explainer) Explain(taskNames ...string) (err error) {
	for _, taskName := range taskNames {
		if _, exist := e.project.Tasks[taskName]; !exist {
			return fmt.Errorf("task %s does not exist", taskName)
		}
		e.explainTask(taskName)
	}
	return nil
}

func (e *Explainer) explainTask(taskName string) {
	task := e.project.Tasks[taskName]
	indentation := strings.Repeat(" ", 18) + strings.Repeat(e.d.Empty, 2)
	parentTasks := task.Extends
	if len(parentTasks) == 0 && task.Extend != "" {
		parentTasks = []string{task.Extend}
	}
	start, check, taskType := e.getCmdPatterns(task)
	e.printField("TASK NAME    ", taskName, indentation)
	e.printField("LOCATION     ", task.GetFileLocation(), indentation)
	e.printField("DESCRIPTION  ", task.Description, indentation)
	e.printField("TASK TYPE    ", taskType, indentation)
	e.printField("PARENT TASKS ", e.listToStr(parentTasks), indentation)
	e.printField("DEPENDENCIES ", e.listToStr(task.Dependencies), indentation)
	e.printField("START        ", e.listToMultiLineStr(start), indentation)
	e.printField("CHECK        ", e.listToMultiLineStr(check), indentation)
	e.printField("INPUTS       ", e.getInputString(task), indentation)
	e.printField("CONFIG       ", e.getConfigString(task), indentation)
	e.printField("ENVIRONMENTS ", e.getEnvString(task), indentation)
}

func (e *Explainer) getCmdPatterns(task *config.Task) (start []string, check []string, taskType string) {
	start, startExist, _ := task.GetStartCmdPatterns()
	check, checkExist, _ := task.GetCheckCmdPatterns()
	if startExist && checkExist {
		return start, check, "Service Task"
	}
	if startExist {
		return start, check, "Command Task"
	}
	return start, check, "Wrapper Task"
}

func (e *Explainer) getInputString(task *config.Task) (inputString string) {
	inputNames := task.Inputs
	inputCount := len(inputNames)
	if inputCount == 0 {
		return ""
	}
	subFieldIndentation := strings.Repeat(" ", 2)
	for _, inputName := range inputNames {
		input := e.project.Inputs[inputName]
		rawInputLines := []string{
			inputName,
			e.getSubFieldString("DESCRIPTION", input.Description, subFieldIndentation),
			e.getSubFieldString("PROMPT     ", input.Prompt, subFieldIndentation),
			e.getSubFieldString("OPTIONS    ", e.listToStr(input.Options), subFieldIndentation),
			e.getSubFieldString("DEFAULT    ", input.DefaultValue, subFieldIndentation),
			e.getSubFieldString("VALIDATION ", input.Validation, subFieldIndentation),
		}
		inputLines := []string{}
		for _, inputLine := range rawInputLines {
			if inputLine != "" {
				inputLines = append(inputLines, inputLine)
			}
		}
		inputString += strings.Trim(strings.Join(inputLines, "\n"), "\n") + "\n"
	}
	return inputString
}

func (e *Explainer) getEnvString(task *config.Task) (envString string) {
	subFieldIndentation := strings.Repeat(" ", 2)
	keys := task.GetEnvKeys()
	sort.Strings(keys)
	for _, key := range keys {
		env, _ := task.GetEnvObject(key)
		rawEnvLines := []string{
			key,
			e.getSubFieldString("FROM   ", env.From, subFieldIndentation),
			e.getSubFieldString("DEFAULT", env.Default, subFieldIndentation),
		}
		envLines := []string{}
		for _, envLine := range rawEnvLines {
			if envLine != "" {
				envLines = append(envLines, envLine)
			}
		}
		envString += strings.Trim(strings.Join(envLines, "\n"), "\n") + "\n"
	}
	return envString
}

func (e *Explainer) getConfigString(task *config.Task) (configStr string) {
	keys := task.GetConfigKeys()
	sort.Strings(keys)
	fieldKeys := e.getFieldKeys(keys)
	lines := []string{}
	for index, key := range keys {
		fieldKey := fieldKeys[index]
		val, _ := task.GetConfigPattern(key)
		fieldVal := e.getStrOrBlank(val)
		lines = append(lines, e.getSubFieldString(fieldKey, fieldVal, ""))
	}
	return strings.Join(lines, "\n")
}

func (e *Explainer) getSubFieldString(subFieldName string, subFieldValue string, subFieldIndentation string) (subFieldStr string) {
	subFieldValue = strings.Trim(subFieldValue, "\n")
	if subFieldValue == "" {
		return ""
	}
	subFieldValueIndentation := subFieldIndentation + strings.Repeat(" ", len(subFieldName)+1)
	subFieldLines := strings.Split(strings.Trim(subFieldValue, "\n"), "\n")
	subFieldValueStr := strings.Join(subFieldLines, "\n  "+subFieldValueIndentation)
	return fmt.Sprintf("%s%s%s :%s %s", subFieldIndentation, e.d.Yellow, subFieldName, e.d.Normal, subFieldValueStr)
}

func (e *Explainer) printField(fieldName string, value string, indentation string) {
	trimmedValue := strings.TrimRight(value, "\n ")
	if trimmedValue == "" {
		return
	}
	indentedValue := str.Indent(trimmedValue, indentation)
	e.logger.DPrintf("%s%s :%s %s\n", e.d.Yellow, fieldName, e.d.Normal, indentedValue)
}

func (e *Explainer) GetZarubaCommand(taskNames []string, autoTerminate bool, autoTerminateDelayInterval string) (command string) {
	command = fmt.Sprintf("zaruba please %s", strings.Join(taskNames, " "))
	inputArgs := []string{}
	for _, taskName := range taskNames {
		task := e.project.Tasks[taskName]
		for _, inputName := range task.Inputs {
			inputValue := ""
			if input := e.project.Inputs[inputName]; input.Secret {
				inputValue = "[HIDDEN_VALUE]"
			} else {
				inputValue = e.project.GetValue(inputName)
			}
			inputArgs = append(inputArgs, fmt.Sprintf("%s=%s", inputName, str.DoubleQuoteShellValue(inputValue)))
		}
	}
	if len(inputArgs) != 0 {
		command = fmt.Sprintf("%s %s", command, strings.Join(inputArgs, " "))
	}
	if autoTerminate {
		command = fmt.Sprintf("%s -t -w %s", command, autoTerminateDelayInterval)
	}
	return command
}
