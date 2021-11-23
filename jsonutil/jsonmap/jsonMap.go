package jsonmap

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/state-alchemists/zaruba/dictutil"
	jsonHelper "github.com/state-alchemists/zaruba/jsonutil/helper"
	"github.com/state-alchemists/zaruba/strutil"
)

type JsonMap struct {
	str *strutil.StrUtil
}

func NewJsonMap(strUtil *strutil.StrUtil) *JsonMap {
	return &JsonMap{
		str: strUtil,
	}
}

func (jsonMap *JsonMap) Validate(mapString string) (valid bool) {
	_, err := jsonHelper.ToDict(mapString)
	return err == nil
}

func (jsonMap *JsonMap) GetValue(mapString, key string) (data interface{}, err error) {
	dict, err := jsonHelper.ToDict(mapString)
	if err != nil {
		return nil, err
	}
	return dict[key], nil
}

func (jsonMap *JsonMap) GetKeys(mapString string) (keys []string, err error) {
	dict, err := jsonHelper.ToDict(mapString)
	if err != nil {
		return nil, err
	}
	return dictutil.DictGetSortedKeys(dict)
}

func (jsonMap *JsonMap) Merge(mapStrings ...string) (mergedMapString string, err error) {
	mergedDict := jsonHelper.Dict{}
	for _, mapString := range mapStrings {
		dict, err := jsonHelper.ToDict(mapString)
		if err != nil {
			return "{}", err
		}
		for key, val := range dict {
			if _, exist := mergedDict[key]; !exist {
				mergedDict[key] = val
			}
		}
	}
	mergedMapBytes, err := json.Marshal(mergedDict)
	if err != nil {
		return "{}", err
	}
	return string(mergedMapBytes), nil
}

func (jsonMap *JsonMap) Set(mapString string, args ...string) (newMapString string, err error) {
	dict, err := jsonHelper.ToDict(mapString)
	if err != nil {
		return mapString, err
	}
	if len(args)%2 != 0 {
		return mapString, fmt.Errorf("invalid number of arguments")
	}
	for index := 0; index < len(args); index += 2 {
		newKey := args[index]
		newVal := jsonHelper.ToInterface(args[index+1])
		dict[newKey] = newVal
	}
	newMapBytes, err := json.Marshal(dict)
	if err != nil {
		return "{}", err
	}
	return string(newMapBytes), nil
}

func (jsonMap *JsonMap) TransformKeys(mapString string, prefix string, suffix string) (newMapString string, err error) {
	dict, err := jsonHelper.ToDict(mapString)
	if err != nil {
		return mapString, err
	}
	newDict := jsonHelper.Dict{}
	for key, val := range dict {
		newKey := fmt.Sprintf("%s%s%s", prefix, key, suffix)
		newDict[newKey] = val
	}
	newMapBytes, err := json.Marshal(newDict)
	if err != nil {
		return "{}", err
	}
	return string(newMapBytes), nil
}

func (jsonMap *JsonMap) CascadePrefixKeys(mapString string, prefix string) (newMapString string, err error) {
	dict, err := jsonHelper.ToDict(mapString)
	if err != nil {
		return mapString, err
	}
	newDict := jsonHelper.Dict{}
	for key, val := range dict {
		newDict[key] = val
		prefixedKeyParts := strings.SplitN(key, "_", 2)
		if len(prefixedKeyParts) < 2 {
			continue
		}
		keyPrefix, key := prefixedKeyParts[0], prefixedKeyParts[1]
		if keyPrefix != prefix {
			continue
		}
		newDict[key] = val
	}
	newMapBytes, err := json.Marshal(newDict)
	if err != nil {
		return "{}", err
	}
	return string(newMapBytes), nil
}

func (jsonMap *JsonMap) GetFromEnv() (mapString string, err error) {
	stringDict := jsonHelper.StringDict{}
	for _, pair := range os.Environ() {
		pairParts := strings.SplitN(pair, "=", 2)
		key, val := pairParts[0], pairParts[1]
		stringDict[key] = val
	}
	dictBytes, err := json.Marshal(stringDict)
	if err != nil {
		return "{}", err
	}
	return string(dictBytes), nil
}

func (jsonMap *JsonMap) ToStringMap(mapString string) (newMapString string, err error) {
	stringDict, err := jsonHelper.ToStringDict(mapString)
	if err != nil {
		return "{}", err
	}
	newMapBytes, err := json.Marshal(stringDict)
	if err != nil {
		return "{}", err
	}
	return string(newMapBytes), nil
}

func (jsonMap *JsonMap) ToVariedStringMap(mapString string, keys ...string) (newMapString string, err error) {
	variedStringDict, err := jsonHelper.ToStringDict(mapString)
	if err != nil {
		return "{}", err
	}
	if len(keys) == 0 {
		for key := range variedStringDict {
			keys = append(keys, key)
		}
	}
	strTransformators := []func(string) string{jsonMap.str.ToKebab, jsonMap.str.ToCamel, jsonMap.str.ToSnake, jsonMap.str.ToLower, jsonMap.str.ToUpper, jsonMap.str.ToUpperSnake, jsonMap.str.DoubleQuote, jsonMap.str.SingleQuote}
	for _, key := range keys {
		val := variedStringDict[key]
		for _, strTransformator := range strTransformators {
			newKey := strTransformator(key)
			_, newKeyAlreadyExist := variedStringDict[newKey]
			if !newKeyAlreadyExist {
				newVal := strTransformator(val)
				variedStringDict[newKey] = newVal
			}
		}
	}
	newMapBytes, err := json.Marshal(variedStringDict)
	if err != nil {
		return "{}", err
	}
	return string(newMapBytes), nil
}

func (jsonMap *JsonMap) ToEnvString(mapString string) (envString string, err error) {
	stringDict, err := jsonHelper.ToStringDict(mapString)
	if err != nil {
		return "", err
	}
	return godotenv.Marshal(stringDict)
}

func (jsonMap *JsonMap) Replace(str string, replacementMapString string) (newStr string, err error) {
	replacementMap, err := jsonHelper.ToStringDict(replacementMapString)
	if err != nil {
		return str, err
	}
	return strutil.StrReplace(str, replacementMap), nil
}
