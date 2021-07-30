package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"github.com/google/uuid"
	"github.com/state-alchemists/zaruba/boolean"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/str"
	yaml "gopkg.in/yaml.v3"
)

type TaskData struct {
	task         *Task
	Zaruba       string
	Name         string
	ProjectName  string
	WorkPath     string
	DirPath      string
	FileLocation string
	Decoration   *output.Decoration
}

func NewTaskData(task *Task) (td *TaskData) {
	nextTask := *task
	nextTask.currentRecursiveLevel++
	return &TaskData{
		task:         &nextTask,
		Zaruba:       "\"${ZARUBA_HOME}/zaruba\"",
		Name:         task.GetName(),
		ProjectName:  task.Project.GetName(),
		WorkPath:     task.GetWorkPath(),
		DirPath:      filepath.Dir(task.GetFileLocation()),
		FileLocation: task.GetFileLocation(),
		Decoration:   task.Project.decoration,
	}
}

func (td *TaskData) GetWorkPath(path string) (absPath string) {
	return td.getAbsPath(td.WorkPath, path)
}

func (td *TaskData) GetRelativePath(path string) (absPath string) {
	return td.getAbsPath(td.DirPath, path)
}

func (td *TaskData) GetConfig(keys ...string) (val string, err error) {
	return td.task.GetConfig(keys...)
}

func (td *TaskData) GetSubConfigKeys(parentKeys ...string) (subKeys []string) {
	configKeys := td.task.GetConfigKeys()
	return str.GetSubKeys(configKeys, parentKeys)
}

func (td *TaskData) GetValue(keys ...string) (val string, err error) {
	return td.task.GetValue(keys...)
}

func (td *TaskData) GetSubValueKeys(parentKeys ...string) (subKeys []string) {
	valueKeys := td.task.GetValueKeys()
	return str.GetSubKeys(valueKeys, parentKeys)
}

func (td *TaskData) GetEnv(key string) (val string, err error) {
	return td.task.GetEnv(key)
}

func (td *TaskData) GetEnvs() (parsedEnv map[string]string, err error) {
	return td.task.GetEnvs()
}

func (td *TaskData) IsTrue(str string) (isTrue bool) {
	return boolean.IsTrue(str)
}

func (td *TaskData) IsFalse(str string) (isFalse bool) {
	return boolean.IsFalse(str)
}

func (td *TaskData) ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func (td *TaskData) EscapeShellArg(s string) (result string) {
	return str.EscapeShellArg(s)
}

func (td *TaskData) Indent(multiLineStr string, indentation string) (result string) {
	return str.Indent(multiLineStr, indentation)
}

func (td *TaskData) GetNewUUID() string {
	return uuid.NewString()
}

func (td *TaskData) Split(s, sep string) []string {
	return strings.Split(s, sep)
}

func (td *TaskData) Join(sep string, a []string) (string, error) {
	return strings.Join(a, sep), nil
}

func (td *TaskData) Trim(str, cutset string) (trimmedStr string) {
	return strings.Trim(str, cutset)
}

func (td *TaskData) ParseJSON(s string) (interface{}, error) {
	if s == "" {
		return make([]interface{}, 0), nil
	}
	var data interface{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (td *TaskData) ParseYAML(s string) (interface{}, error) {
	if s == "" {
		return make([]interface{}, 0), nil
	}
	var data interface{}
	if err := yaml.Unmarshal([]byte(s), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (td *TaskData) ReadFile(filePath string) (fileContent string, err error) {
	absFilePath := td.GetWorkPath(filePath)
	return file.ReadText(absFilePath)
}

func (td *TaskData) ListDir(dirPath string) (fileNames []string, err error) {
	absDirPath := td.GetWorkPath(dirPath)
	return file.ListDir(absDirPath)
}

func (td *TaskData) ParseFile(filePath string) (parsedStr string, err error) {
	absFilePath := td.GetWorkPath(filePath)
	pattern, err := td.ReadFile(absFilePath)
	if err != nil {
		return "", err
	}
	templateName := fmt.Sprintf("File: %s", absFilePath)
	tmpl, err := template.New(templateName).Option("missingkey=zero").Parse(pattern)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err = tmpl.Execute(&b, td); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (td *TaskData) WriteFile(filePath string, content string) (err error) {
	return file.WriteText(filePath, content, 0755)
}

func (td *TaskData) Add(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() + bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() + int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) + bv.Float(), nil
		default:
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) + bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() + bv.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) + bv.Float(), nil
		default:
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() + float64(bv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() + float64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() + bv.Float(), nil
		default:
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	default:
		return nil, fmt.Errorf("add: unknown type for %q (%T)", av, a)
	}
}

func (td *TaskData) Subtract(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() - bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() - int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) - bv.Float(), nil
		default:
			return nil, fmt.Errorf("subtract: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) - bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() - bv.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) - bv.Float(), nil
		default:
			return nil, fmt.Errorf("subtract: unknown type for %q (%T)", bv, b)
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() - float64(bv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() - float64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() - bv.Float(), nil
		default:
			return nil, fmt.Errorf("subtract: unknown type for %q (%T)", bv, b)
		}
	default:
		return nil, fmt.Errorf("subtract: unknown type for %q (%T)", av, a)
	}
}

func (td *TaskData) Multiply(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() * bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() * int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) * bv.Float(), nil
		default:
			return nil, fmt.Errorf("multiply: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) * bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() * bv.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) * bv.Float(), nil
		default:
			return nil, fmt.Errorf("multiply: unknown type for %q (%T)", bv, b)
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() * float64(bv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() * float64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() * bv.Float(), nil
		default:
			return nil, fmt.Errorf("multiply: unknown type for %q (%T)", bv, b)
		}
	default:
		return nil, fmt.Errorf("multiply: unknown type for %q (%T)", av, a)
	}
}

func (td *TaskData) Divide(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() / bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() / int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) / bv.Float(), nil
		default:
			return nil, fmt.Errorf("divide: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) / bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() / bv.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) / bv.Float(), nil
		default:
			return nil, fmt.Errorf("divide: unknown type for %q (%T)", bv, b)
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() / float64(bv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() / float64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() / bv.Float(), nil
		default:
			return nil, fmt.Errorf("divide: unknown type for %q (%T)", bv, b)
		}
	default:
		return nil, fmt.Errorf("divide: unknown type for %q (%T)", av, a)
	}
}

func (td *TaskData) In(l, v interface{}) (bool, error) {
	lv := reflect.ValueOf(l)
	vv := reflect.ValueOf(v)

	switch lv.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < lv.Len(); i++ {
			lvv := lv.Index(i)
			switch lvv.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				switch vv.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if vv.Int() == lvv.Int() {
						return true, nil
					}
				}
			case reflect.Float32, reflect.Float64:
				switch vv.Kind() {
				case reflect.Float32, reflect.Float64:
					if vv.Float() == lvv.Float() {
						return true, nil
					}
				}
			case reflect.String:
				if vv.Type() == lvv.Type() && vv.String() == lvv.String() {
					return true, nil
				}
			}
		}
	case reflect.String:
		if vv.Type() == lv.Type() && strings.Contains(lv.String(), vv.String()) {
			return true, nil
		}
	}

	return false, nil
}

func (td *TaskData) getAbsPath(parentPath, path string) (absPath string) {
	if filepath.IsAbs(path) {
		return path
	}
	absParentPath, _ := filepath.Abs(parentPath)
	return filepath.Join(absParentPath, path)
}
