package jsonlist

import (
	"encoding/json"
	"fmt"
	"strings"

	jsonHelper "github.com/state-alchemists/zaruba/jsonutil/helper"
)

type List []interface{}
type StringList []string

type JsonList struct{}

func NewJsonList() *JsonList {
	return &JsonList{}
}

func (jsonList *JsonList) listToLines(list List) (lines StringList) {
	lines = StringList{}
	for _, val := range list {
		lines = append(lines, jsonHelper.FromInterface(val))
	}
	return lines
}

func (jsonList *JsonList) GetList(listString string) (list List, err error) {
	list = List{}
	err = json.Unmarshal([]byte(listString), &list)
	return list, err
}

func (jsonList *JsonList) GetStringList(listString string) (stringList StringList, err error) {
	stringList = StringList{}
	list, err := jsonList.GetList(listString)
	if err != nil {
		return stringList, err
	}
	stringList = jsonList.listToLines(list)
	return stringList, err
}

func (jsonList *JsonList) Validate(listString string) (valid bool) {
	_, err := jsonList.GetList(listString)
	return err == nil
}

func (jsonList *JsonList) GetValue(listString string, index int) (data interface{}, err error) {
	list, err := jsonList.GetList(listString)
	if err != nil {
		return nil, err
	}
	if index < 0 || index > len(list) {
		return nil, fmt.Errorf("index of bound")
	}
	return list[index], nil
}

func (jsonList *JsonList) GetLength(listString string) (length int, err error) {
	list, err := jsonList.GetList(listString)
	if err != nil {
		return -1, err
	}
	return len(list), nil
}

func (jsonList *JsonList) Append(listString string, value string) (newListString string, err error) {
	list, err := jsonList.GetList(listString)
	if err != nil {
		return "[]", err
	}
	list = append(list, jsonHelper.ToInterface(value))
	newListBytes, err := json.Marshal(list)
	if err != nil {
		return "[]", err
	}
	return string(newListBytes), nil
}

func (jsonList *JsonList) Set(listString string, index int, value string) (newListString string, err error) {
	list, err := jsonList.GetList(listString)
	if err != nil {
		return "[]", err
	}
	if index < 0 || index > len(list) {
		return "[]", fmt.Errorf("index of bound")
	}
	list[index] = jsonHelper.ToInterface(value)
	newListBytes, err := json.Marshal(list)
	if err != nil {
		return "[]", err
	}
	return string(newListBytes), nil
}

func (jsonList *JsonList) Merge(listStrings ...string) (mergedListString string, err error) {
	mergedList := List{}
	for _, listString := range listStrings {
		list, err := jsonList.GetList(listString)
		if err != nil {
			return "[]", err
		}
		mergedList = append(mergedList, list...)
	}
	mergedListBytes, err := json.Marshal(mergedList)
	if err != nil {
		return "{}", err
	}
	return string(mergedListBytes), nil
}

func (jsonList *JsonList) Join(listString string, separator string) (str string, err error) {
	lines, err := jsonList.GetStringList(listString)
	if err != nil {
		return "", err
	}
	str = strings.Join(lines, separator)
	return str, nil
}

func (jsonList *JsonList) GetIndex(listString string, elementString string) (index int, err error) {
	list, err := jsonList.GetStringList(listString)
	if err != nil {
		return -1, err
	}
	for index, existingRow := range list {
		if existingRow == elementString {
			return index, nil
		}
	}
	return -1, nil
}

func (jsonList *JsonList) Contain(listString string, elementString string) (exist bool, err error) {
	index, err := jsonList.GetIndex(listString, elementString)
	if err != nil {
		return false, err
	}
	return index > -1, nil
}
