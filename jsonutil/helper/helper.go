package common

import (
	"encoding/json"
	"fmt"
)

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
