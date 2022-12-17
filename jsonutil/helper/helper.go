package common

import (
	"encoding/json"
	"fmt"
)

type Dict map[string]interface{}
type StringDict map[string]string
type List []interface{}
type StringList []string

func NormalizeObj(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = NormalizeObj(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = NormalizeObj(v)
		}
	}
	return i
}

func FromInterface(obj interface{}) (jsonAny string) {
	if obj == nil {
		return ""
	}
	jsonByte, err := json.Marshal(obj)
	if err == nil && len(jsonByte) > 0 {
		ch := jsonByte[0]
		if ch == []byte("{")[0] || ch == []byte("[")[0] {
			return string(jsonByte)
		}
	}
	return fmt.Sprintf("%v", obj)
}

func ToInterface(jsonAny string) (obj interface{}) {
	if err := json.Unmarshal([]byte(jsonAny), &obj); err != nil {
		return jsonAny
	}
	return obj
}

func FromDict(dict Dict) (jsonMap string, err error) {
	jsonByte, err := json.Marshal(dict)
	if err != nil {
		return "{}", err
	}
	return string(jsonByte), err
}

func ToDict(jsonMap string) (dict Dict, err error) {
	dict = Dict{}
	err = json.Unmarshal([]byte(jsonMap), &dict)
	return dict, err
}

func FromStringDict(stringDict StringDict) (jsonMap string, err error) {
	jsonByte, err := json.Marshal(stringDict)
	if err != nil {
		return "{}", err
	}
	return string(jsonByte), err
}

func ToStringDict(jsonMap string) (stringDict StringDict, err error) {
	stringDict = StringDict{}
	err = json.Unmarshal([]byte(jsonMap), &stringDict)
	if err == nil {
		return stringDict, nil
	}
	dict, err := ToDict(jsonMap)
	if err != nil {
		return stringDict, err
	}
	stringDict = StringDict{}
	for key, val := range dict {
		stringDict[key] = FromInterface(val)
	}
	return stringDict, err
}

func FromList(list List) (jsonList string, err error) {
	jsonByte, err := json.Marshal(list)
	if err != nil {
		return "[]", err
	}
	return string(jsonByte), err
}

func ToList(jsonList string) (list List, err error) {
	list = List{}
	err = json.Unmarshal([]byte(jsonList), &list)
	return list, err
}

func FromStringList(stringList StringList) (jsonList string, err error) {
	jsonByte, err := json.Marshal(stringList)
	if err != nil {
		return "[]", err
	}
	return string(jsonByte), err
}

func ToStringList(jsonString string) (jsonList StringList, err error) {
	jsonList = StringList{}
	err = json.Unmarshal([]byte(jsonString), &jsonList)
	if err == nil {
		return jsonList, nil
	}
	list, err := ToList(jsonString)
	if err != nil {
		return jsonList, err
	}
	jsonList = StringList{}
	for _, val := range list {
		jsonList = append(jsonList, FromInterface(val))
	}
	return jsonList, err
}
