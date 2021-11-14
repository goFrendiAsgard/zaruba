package core

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type TaskEnvKeyCheckData struct {
	Envs map[string]map[string]interface{} `yaml:"envs,omitempty"`
}

type TaskWithEnvKeyCheckData struct {
	Tasks map[string]TaskEnvKeyCheckData `yaml:"tasks,omitempty"`
}

type TaskKeyCheckData struct {
	Tasks map[string]map[string]interface{} `yaml:"tasks,omitempty"`
}

type EnvKeyCheckData struct {
	Envs map[string]map[string]map[string]interface{} `yaml:"envs,omitempty"`
}

type InputKeyCheckData struct {
	Variables map[string]map[string]interface{} `yaml:"inputs,omitempty"`
}

// KeyValidator structure validator
type KeyValidator struct {
	fileName                string
	rawData                 map[string]interface{}
	envKeyCheckData         EnvKeyCheckData
	taskKeyCheckData        TaskKeyCheckData
	taskWithEnvKeyCheckData TaskWithEnvKeyCheckData
	inputKeyCheckData       InputKeyCheckData
}

func NewKeyValidator(fileName string) (sv *KeyValidator) {
	return &KeyValidator{
		fileName: fileName,
		rawData:  map[string]interface{}{},
		envKeyCheckData: EnvKeyCheckData{
			Envs: map[string]map[string]map[string]interface{}{},
		},
		inputKeyCheckData: InputKeyCheckData{
			Variables: map[string]map[string]interface{}{},
		},
		taskKeyCheckData: TaskKeyCheckData{
			Tasks: map[string]map[string]interface{}{},
		},
		taskWithEnvKeyCheckData: TaskWithEnvKeyCheckData{
			Tasks: map[string]TaskEnvKeyCheckData{},
		},
	}
}

func (kv *KeyValidator) Validate() (b []byte, err error) {
	b, err = ioutil.ReadFile(kv.fileName)
	if err != nil {
		return b, fmt.Errorf("error reading file '%s': %s", kv.fileName, err)
	}
	if err = yaml.Unmarshal(b, &kv.rawData); err != nil {
		return b, fmt.Errorf("error parsing YAML '%s': %s", kv.fileName, err)
	}
	if err = yaml.Unmarshal(b, &kv.envKeyCheckData); err != nil {
		return b, fmt.Errorf("error parsing YAML '%s': %s", kv.fileName, err)
	}
	if err = yaml.Unmarshal(b, &kv.inputKeyCheckData); err != nil {
		return b, fmt.Errorf("error parsing YAML '%s': %s", kv.fileName, err)
	}
	if err = yaml.Unmarshal(b, &kv.taskKeyCheckData); err != nil {
		return b, fmt.Errorf("error parsing YAML '%s': %s", kv.fileName, err)
	}
	if err = yaml.Unmarshal(b, &kv.taskWithEnvKeyCheckData); err != nil {
		return b, fmt.Errorf("error parsing YAML '%s': %s", kv.fileName, err)
	}
	if err = kv.checkProjectValidKeys(); err != nil {
		return b, err
	}
	if err = kv.checkEnvValidKeys(); err != nil {
		return b, err
	}
	if err = kv.checkInputValidKeys(); err != nil {
		return b, err
	}
	if err = kv.checkTaskValidKeys(); err != nil {
		return b, err
	}
	if err = kv.checkTaskEnvValidKeys(); err != nil {
		return b, err
	}
	return b, nil
}

func (kv *KeyValidator) checkProjectValidKeys() (err error) {
	validKeys := []string{"includes", "tasks", "name", "inputs", "envs", "configs"}
	for key := range kv.rawData {
		isValid := false
		for _, validKey := range validKeys {
			if validKey == key {
				isValid = true
				break
			}
		}
		if !isValid {
			return fmt.Errorf("invalid key on '%s': %s", kv.fileName, key)
		}
	}
	return nil
}

func (kv *KeyValidator) checkTaskValidKeys() (err error) {
	validKeys := []string{"start", "check", "timeout", "private", "autoTerminate", "extend", "extends", "location", "configRef", "configRefs", "configs", "envRef", "envRefs", "envs", "dependencies", "inputs", "description", "icon", "saveLog", "syncEnv", "syncEnvLocation"}
	for taskName, task := range kv.taskKeyCheckData.Tasks {
		for key := range task {
			valid := false
			for _, validKey := range validKeys {
				if key == validKey {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("invalid key on '%s': tasks[%s][%s]", kv.fileName, taskName, key)
			}
		}
	}
	return nil
}

func (kv *KeyValidator) checkEnvValidKeys() (err error) {
	for envRefName, env := range kv.envKeyCheckData.Envs {
		if err = kv.checkEnvMapValidKeys(env, fmt.Sprintf("envs[%s]", envRefName)); err != nil {
			return err
		}
	}
	return nil
}

func (kv *KeyValidator) checkInputValidKeys() (err error) {
	validKeys := []string{"default", "description", "secret", "validation", "options", "prompt", "allowCustom"}
	for inputName, input := range kv.inputKeyCheckData.Variables {
		for key := range input {
			valid := false
			for _, validKey := range validKeys {
				if key == validKey {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("invalid key on '%s': inputs[%s][%s]", kv.fileName, inputName, key)
			}
		}
	}
	return nil
}

func (kv *KeyValidator) checkTaskEnvValidKeys() (err error) {
	for taskName, task := range kv.taskWithEnvKeyCheckData.Tasks {
		if err = kv.checkEnvMapValidKeys(task.Envs, fmt.Sprintf("tasks[%s][envs]", taskName)); err != nil {
			return err
		}
	}
	return nil
}

func (kv *KeyValidator) checkEnvMapValidKeys(envMap map[string]map[string]interface{}, errorKeyPrefix string) (err error) {
	validKeys := []string{"from", "default"}
	for envName := range envMap {
		for key := range envMap[envName] {
			valid := false
			for _, validKey := range validKeys {
				if key == validKey {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("invalid key on '%s': %s[%s][%s]", kv.fileName, errorKeyPrefix, envName, key)
			}
		}
	}
	return nil
}
