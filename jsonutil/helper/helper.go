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

func FromInterface(obj interface{}) (jsonString string) {
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

func ToInterface(jsonString string) (obj interface{}) {
	if err := json.Unmarshal([]byte(jsonString), &obj); err != nil {
		return jsonString
	}
	return obj
}

func FromDict(dict Dict) (jsonString string, err error) {
	jsonByte, err := json.Marshal(dict)
	if err != nil {
		return "{}", err
	}
	return string(jsonByte), err
}

func ToDict(jsonString string) (dict Dict, err error) {
	dict = Dict{}
	err = json.Unmarshal([]byte(jsonString), &dict)
	return dict, err
}

func FromStringDict(stringDict StringDict) (jsonString string, err error) {
	jsonByte, err := json.Marshal(stringDict)
	if err != nil {
		return "{}", err
	}
	return string(jsonByte), err
}

func ToStringDict(jsonString string) (stringDict StringDict, err error) {
	stringDict = StringDict{}
	err = json.Unmarshal([]byte(jsonString), &stringDict)
	if err == nil {
		return stringDict, nil
	}
	dict, err := ToDict(jsonString)
	if err != nil {
		return stringDict, err
	}
	stringDict = StringDict{}
	for key, val := range dict {
		stringDict[key] = FromInterface(val)
	}
	return stringDict, err
}

func FromList(list List) (jsonString string, err error) {
	jsonByte, err := json.Marshal(list)
	if err != nil {
		return "[]", err
	}
	return string(jsonByte), err
}

func ToList(jsonString string) (list List, err error) {
	list = List{}
	err = json.Unmarshal([]byte(jsonString), &list)
	return list, err
}

func FromStringList(stringList StringList) (jsonString string, err error) {
	jsonByte, err := json.Marshal(stringList)
	if err != nil {
		return "[]", err
	}
	return string(jsonByte), err
}

func ToStringList(jsonString string) (stringList StringList, err error) {
	stringList = StringList{}
	err = json.Unmarshal([]byte(jsonString), &stringList)
	if err == nil {
		return stringList, nil
	}
	list, err := ToList(jsonString)
	if err != nil {
		return stringList, err
	}
	stringList = StringList{}
	for _, val := range list {
		stringList = append(stringList, FromInterface(val))
	}
	return stringList, err
}
