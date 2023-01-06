package helper

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Set(jsonStr string, key string, jsonStrVal string) (newJsonStr string, err error) {
	var val interface{} = getValFromJsonStrVal(jsonStrVal)
	keys := strings.Split(key, ".")
	newJsonStr, err = set(jsonStr, keys, val)
	if err != nil {
		err = fmt.Errorf("cannot fetch key %s from %s: %#v", key, jsonStr, err)
	}
	return newJsonStr, err

}

func getValFromJsonStrVal(jsonStrVal string) (val interface{}) {
	val = nil
	if err := json.Unmarshal([]byte(jsonStrVal), &val); err != nil {
		val = jsonStrVal
	}
	return val
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
