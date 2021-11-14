package jsonutil

import (
	"encoding/json"
	"strings"

	"github.com/state-alchemists/zaruba/yamlstyler"
	"gopkg.in/yaml.v3"
)

type JsonUtil struct{}

func NewJsonUtil() *JsonUtil {
	return &JsonUtil{}
}

func (jsonUtil *JsonUtil) normalizeObj(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = jsonUtil.normalizeObj(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = jsonUtil.normalizeObj(v)
		}
	}
	return i
}

func (jsonUtil *JsonUtil) FromYaml(yamlString string) (jsonString string, err error) {
	var interfaceContent interface{}
	if err = yaml.Unmarshal([]byte(yamlString), &interfaceContent); err != nil {
		return "", err
	}
	interfaceContent = jsonUtil.normalizeObj(interfaceContent)
	resultB, err := json.Marshal(interfaceContent)
	if err != nil {
		return "", err
	}
	return string(resultB), nil
}

func (jsonUtil *JsonUtil) ToYaml(jsonString string) (yamlString string, err error) {
	var interfaceContent interface{}
	if err := json.Unmarshal([]byte(jsonString), &interfaceContent); err != nil {
		return "", err
	}
	interfaceContent = jsonUtil.normalizeObj(interfaceContent)
	yamlContentB, err := yaml.Marshal(interfaceContent)
	if err != nil {
		return "", err
	}
	yamlContent := string(yamlContentB)
	yamlLines := strings.Split(yamlContent, "\n")
	for _, styler := range []yamlstyler.YamlStyler{yamlstyler.TwoSpaces, yamlstyler.FixEmoji} {
		yamlLines = styler(yamlLines)
	}
	return strings.Join(yamlLines, "\n"), nil
}
