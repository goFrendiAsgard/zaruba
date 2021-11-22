package jsonutil

import (
	"encoding/json"
	"strings"

	jsonHelper "github.com/state-alchemists/zaruba/jsonutil/helper"
	"github.com/state-alchemists/zaruba/jsonutil/jsonlist"
	"github.com/state-alchemists/zaruba/jsonutil/jsonmap"
	"github.com/state-alchemists/zaruba/strutil"
	"github.com/state-alchemists/zaruba/yamlstyler"
	"gopkg.in/yaml.v3"
)

type JsonUtil struct {
	Map  *jsonmap.JsonMap
	List *jsonlist.JsonList
}

func NewJsonUtil(strUtil *strutil.StrUtil) *JsonUtil {
	return &JsonUtil{
		Map:  jsonmap.NewJsonMap(strUtil),
		List: jsonlist.NewJsonList(),
	}
}

func (jsonUtil *JsonUtil) FromInterface(obj interface{}) (jsonString string) {
	return jsonHelper.FromInterface(obj)
}

func (jsonUtil *JsonUtil) ToInterface(jsonString string) (obj interface{}) {
	return jsonHelper.ToInterface(jsonString)
}

func (jsonUtil *JsonUtil) FromYaml(yamlString string) (jsonString string, err error) {
	var interfaceContent interface{}
	if err = yaml.Unmarshal([]byte(yamlString), &interfaceContent); err != nil {
		return "", err
	}
	interfaceContent = jsonHelper.NormalizeObj(interfaceContent)
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
	interfaceContent = jsonHelper.NormalizeObj(interfaceContent)
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

func (jsonUtil *JsonUtil) FromDict(dict jsonHelper.Dict) (jsonString string, err error) {
	return jsonHelper.FromDict(dict)
}

func (jsonUtil *JsonUtil) ToDict(jsonString string) (dict jsonHelper.Dict, err error) {
	return jsonHelper.ToDict(jsonString)
}

func (jsonUtil *JsonUtil) FromStringDict(stringDict jsonHelper.StringDict) (jsonString string, err error) {
	return jsonHelper.FromStringDict(stringDict)
}

func (jsonUtil *JsonUtil) ToStringDict(jsonString string) (stringDict jsonHelper.StringDict, err error) {
	return jsonHelper.ToStringDict(jsonString)
}

func (jsonUtil *JsonUtil) FromList(list jsonHelper.List) (jsonString string, err error) {
	return jsonHelper.FromList(list)
}

func (jsonUtil *JsonUtil) ToList(jsonString string) (list jsonHelper.List, err error) {
	return jsonHelper.ToList(jsonString)
}

func (jsonUtil *JsonUtil) FromStringList(stringList jsonHelper.StringList) (jsonString string, err error) {
	return jsonHelper.FromStringList(stringList)
}

func (jsonUtil *JsonUtil) ToStringList(jsonString string) (stringList jsonHelper.StringList, err error) {
	return jsonHelper.ToStringList(jsonString)
}
