package jsonlist

import (
	"encoding/json"
	"fmt"
	"strings"

	jsonHelper "github.com/state-alchemists/zaruba/jsonutil/helper"
	"github.com/state-alchemists/zaruba/strutil"
)

type JsonListUtil struct{}

func NewJsonListUtil() *JsonListUtil {
	return &JsonListUtil{}
}

func (jsonListUtil *JsonListUtil) Validate(jsonList string) (valid bool) {
	_, err := jsonHelper.ToList(jsonList)
	return err == nil
}

func (jsonListUtil *JsonListUtil) GetValue(jsonList string, index int) (data interface{}, err error) {
	list, err := jsonHelper.ToList(jsonList)
	if err != nil {
		return nil, err
	}
	if index < 0 || index > len(list) {
		return nil, fmt.Errorf("index of bound")
	}
	return list[index], nil
}

func (jsonListUtil *JsonListUtil) GetLength(jsonList string) (length int, err error) {
	list, err := jsonHelper.ToList(jsonList)
	if err != nil {
		return -1, err
	}
	return len(list), nil
}

func (jsonListUtil *JsonListUtil) Append(jsonList string, values ...string) (newJsonList string, err error) {
	list, err := jsonHelper.ToList(jsonList)
	if err != nil {
		return "[]", err
	}
	for _, value := range values {
		list = append(list, jsonHelper.ToInterface(value))
	}
	newListBytes, err := json.Marshal(list)
	if err != nil {
		return "[]", err
	}
	return string(newListBytes), nil
}

func (jsonListUtil *JsonListUtil) Set(jsonList string, index int, value string) (newJsonList string, err error) {
	list, err := jsonHelper.ToList(jsonList)
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

func (jsonListUtil *JsonListUtil) Merge(jsonLists ...string) (jsonListMerged string, err error) {
	mergedList := jsonHelper.List{}
	for _, jsonList := range jsonLists {
		list, err := jsonHelper.ToList(jsonList)
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

func (jsonListUtil *JsonListUtil) Join(jsonStrList string, separator string) (str string, err error) {
	lines, err := jsonHelper.ToStringList(jsonStrList)
	if err != nil {
		return "", err
	}
	str = strings.Join(lines, separator)
	return str, nil
}

func (jsonListUtil *JsonListUtil) GetIndex(jsonStrList string, elementString string) (index int, err error) {
	list, err := jsonHelper.ToStringList(jsonStrList)
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

func (jsonListUtil *JsonListUtil) Contain(jsonStrList string, elementString string) (exist bool, err error) {
	index, err := jsonListUtil.GetIndex(jsonStrList, elementString)
	if err != nil {
		return false, err
	}
	return index > -1, nil
}

func (jsonListutil *JsonListUtil) GetLinesSubmatch(jsonStrList, jsonStrListPattern string, desiredPatternIndex int) (matchIndex int, jsonSubmatch string, err error) {
	lines, patterns, err := jsonListutil.prepareLinesAndPattern(jsonStrList, jsonStrListPattern)
	if err != nil {
		return -1, "[]", err
	}
	matchIndex, submatch, err := strutil.StrGetLineSubmatch(lines, patterns, desiredPatternIndex)
	if err != nil {
		return -1, "[]", err
	}
	jsonSubmatch, err = jsonHelper.FromStringList(submatch)
	return matchIndex, jsonSubmatch, err
}

func (jsonListUtil *JsonListUtil) ReplaceLineAtIndex(jsonStrList string, index int, jsonStrListReplacement string) (newJsonLines string, err error) {
	return jsonListUtil.replaceLineAtIndex(jsonStrList, index, jsonStrListReplacement, "REPLACE")
}

func (jsonListUtil *JsonListUtil) InsertLineAfterIndex(jsonStrList string, index int, jsonStrListReplacement string) (newJsonLines string, err error) {
	return jsonListUtil.replaceLineAtIndex(jsonStrList, index, jsonStrListReplacement, "AFTER")
}

func (jsonListUtil *JsonListUtil) InsertLineBeforeIndex(jsonStrList string, index int, jsonStrListReplacement string) (newJsonLines string, err error) {
	return jsonListUtil.replaceLineAtIndex(jsonStrList, index, jsonStrListReplacement, "BEFORE")
}

func (jsonListUtil *JsonListUtil) replaceLineAtIndex(jsonStrList string, index int, jsonStrListReplacement string, mode string) (newJsonLines string, err error) {
	strList, replacements, err := jsonListUtil.prepareLinesAndPattern(jsonStrList, jsonStrListReplacement)
	if err != nil {
		return "[]", err
	}
	if index < 0 {
		index = len(strList) + index
	}
	switch mode {
	case "BEFORE":
		replacements = append(replacements, strList[index])
	case "AFTER":
		replacements = append([]string{strList[index]}, replacements...)
	case "REPLACE":
	default:
	}
	newLines, err := strutil.StrReplaceLineAtIndex(strList, index, replacements)
	if err != nil {
		return jsonStrList, err
	}
	return jsonHelper.FromStringList(newLines)
}

func (jsonListUtil *JsonListUtil) prepareLinesAndPattern(jsonStrList, jsonStrListPattern string) (strList jsonHelper.StringList, patterns jsonHelper.StringList, err error) {
	strList, err = jsonHelper.ToStringList(jsonStrList)
	if err != nil {
		return strList, patterns, err
	}
	patterns, patternErr := jsonHelper.ToStringList(jsonStrListPattern)
	if patternErr != nil {
		patterns = []string{jsonStrListPattern}
	}
	return strList, patterns, err
}
