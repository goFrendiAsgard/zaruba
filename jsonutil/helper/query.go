package helper

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var listKeyPattern = regexp.MustCompile(`\[([\-]?[0-9]+)\]`)

func isListKey(key string) bool {
	return listKeyPattern.MatchString(key)
}

func getIndexFromKey(key string) int {
	matches := listKeyPattern.FindStringSubmatch(key)
	if len(matches) > 0 {
		index, err := strconv.Atoi(matches[1])
		if err != nil {
			return -1
		}
		return index
	}
	return -1
}

func Get(jsonStr string, key string) (result string, err error) {
	keys := strings.Split(key, ".")
	val, err := get(jsonStr, keys)
	if err != nil {
		return "", err
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

func Set(jsonStr string, key string, jsonStrVal string) (newJsonStr string, err error) {
	var val interface{} = nil
	if err = json.Unmarshal([]byte(jsonStrVal), &val); err != nil {
		return newJsonStr, fmt.Errorf("cannot unmarshal %s: %#v", jsonStrVal, err)
	}
	keys := strings.Split(key, ".")
	newJsonStr, err = set(jsonStr, keys, val)
	if err != nil {
		err = fmt.Errorf("cannot fetch key %s from %s: %#v", key, jsonStr, err)
	}
	return newJsonStr, err

}

func set(jsonStr string, keys []string, val interface{}) (newJsonStr string, err error) {
	if len(keys) == 0 {
		return jsonStr, fmt.Errorf("key not found on %s: %#v", jsonStr, keys)
	}
	newVal, err := setNewVal(jsonStr, keys, val)
	if err != nil {
		return jsonStr, err
	}
	newJsonBytes, err := json.Marshal(newVal)
	if err != nil {
		return jsonStr, err
	}
	return string(newJsonBytes), nil
}

func setNewVal(jsonStr string, keys []string, val interface{}) (newVal interface{}, err error) {
	key, newKeys := keys[0], keys[1:]
	if isListKey(key) {
		return setList(jsonStr, key, newKeys, val)
	}
	return setMap(jsonStr, key, newKeys, val)
}

func setList(jsonStr, key string, keys []string, val interface{}) (newVal interface{}, err error) {
	index := getIndexFromKey(key)
	var jsonList PList = &[]interface{}{}
	if err = json.Unmarshal([]byte(jsonStr), jsonList); err != nil {
		return nil, err
	}
	if index < 0 {
		index = len(*jsonList) + index
	}
	if len(keys) == 0 {
		// this is last element
		(*jsonList)[index] = val
		return *jsonList, err
	}
	subJsonByte, err := json.Marshal((*jsonList)[index])
	if err != nil {
		return nil, err
	}
	newSubVal, err := setNewVal(string(subJsonByte), keys, val)
	if err != nil {
		return nil, err
	}
	(*jsonList)[index] = newSubVal
	return *jsonList, err
}

func setMap(jsonStr, key string, keys []string, val interface{}) (newVal interface{}, err error) {
	var jsonMap PDict = &map[string]interface{}{}
	if err = json.Unmarshal([]byte(jsonStr), jsonMap); err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		// this is last element
		(*jsonMap)[key] = val
		return *jsonMap, err
	}
	subJsonByte, err := json.Marshal((*jsonMap)[key])
	if err != nil {
		return nil, err
	}
	newSubVal, err := setNewVal(string(subJsonByte), keys, val)
	if err != nil {
		return nil, err
	}
	(*jsonMap)[key] = newSubVal
	return *jsonMap, err
}
