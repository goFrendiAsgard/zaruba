package helper

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Get(jsonStr string, key string) (result string, err error) {
	keys := strings.Split(key, ".")
	val, err := get(jsonStr, keys)
	if err != nil {
		return "", err
	}
	result, err = getResultFromVal(val)
	if err != nil {
		return "", err
	}
	return result, err
}

func getResultFromVal(val interface{}) (result string, err error) {
	if valStr, ok := val.(string); ok {
		return valStr, nil
	}
	resultByte, err := json.Marshal(val)
	if err != nil {
		return "", err
	}
	return string(resultByte), err
}

func get(jsonStr string, keys []string) (val interface{}, err error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("key not found on %s: %#v", jsonStr, keys)
	}
	key, newKeys := keys[0], keys[1:]
	newVal, err := getNewVal(key, jsonStr)
	if err != nil {
		return val, err
	}
	// this is the last index
	if len(newKeys) == 0 {
		return newVal, err
	}
	// recursive
	newJsonBytes, err := json.Marshal(newVal)
	if err != nil {
		return val, err
	}
	return get(string(newJsonBytes), newKeys)
}

func getNewVal(key, jsonStr string) (newVal interface{}, err error) {
	if isListKey(key) {
		return getFromList(jsonStr, key)
	}
	return getFromDict(jsonStr, key)
}

func getFromList(jsonStr string, key string) (val interface{}, err error) {
	index := getIndexFromKey(key)
	var jsonList PList = &[]interface{}{}
	if err = json.Unmarshal([]byte(jsonStr), jsonList); err != nil {
		return val, err
	}
	if index < 0 {
		index = len(*jsonList) + index
	}
	// get newVal
	if index >= len(*jsonList) || index < 0 {
		return nil, fmt.Errorf("list index out of bound")
	}
	return (*jsonList)[index], nil
}

func getFromDict(jsonStr string, key string) (val interface{}, err error) {
	var jsonMap PDict = &map[string]interface{}{}
	if err = json.Unmarshal([]byte(jsonStr), jsonMap); err != nil {
		return val, err
	}
	// get newVal
	return (*jsonMap)[key], nil
}
