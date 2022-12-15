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

type JsonMapUtil struct {
	str *strutil.StrUtil
}

func NewJsonMapUtil(strUtil *strutil.StrUtil) *JsonMapUtil {
	return &JsonMapUtil{
		str: strUtil,
	}
}

func (jsonMapUtil *JsonMapUtil) Validate(jsonMap string) (valid bool) {
	_, err := jsonHelper.ToDict(jsonMap)
	return err == nil
}

func (jsonMapUtil *JsonMapUtil) GetValue(jsonMap, key string) (data interface{}, err error) {
	dict, err := jsonHelper.ToDict(jsonMap)
	if err != nil {
		return nil, err
	}
	return dict[key], nil
}

func (jsonMapUtil *JsonMapUtil) GetKeys(jsonMap string) (keys []string, err error) {
	dict, err := jsonHelper.ToDict(jsonMap)
	if err != nil {
		return nil, err
	}
	return dictutil.DictGetSortedKeys(dict)
}

func (jsonMapUtil *JsonMapUtil) Merge(jsonMaps ...string) (jsonMapMerged string, err error) {
	mergedDict := jsonHelper.Dict{}
	for _, mapString := range jsonMaps {
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

func (jsonMapUtil *JsonMapUtil) Set(jsonMap string, args ...string) (newJsonMap string, err error) {
	dict, err := jsonHelper.ToDict(jsonMap)
	if err != nil {
		return jsonMap, err
	}
	if len(args)%2 != 0 {
		return jsonMap, fmt.Errorf("invalid number of arguments")
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

func (jsonMapUtil *JsonMapUtil) TransformKeys(jsonMap string, prefix string, suffix string, transformers ...func(string) string) (newJsonMap string, err error) {
	dict, err := jsonHelper.ToDict(jsonMap)
	if err != nil {
		return jsonMap, err
	}
	newDict := jsonHelper.Dict{}
	for key, val := range dict {
		newKey := key
		for _, transformer := range transformers {
			newKey = transformer(newKey)
		}
		newKey = fmt.Sprintf("%s%s%s", prefix, newKey, suffix)
		newDict[newKey] = val
	}
	newMapBytes, err := json.Marshal(newDict)
	if err != nil {
		return "{}", err
	}
	return string(newMapBytes), nil
}

func (jsonMapUtil *JsonMapUtil) CascadePrefixKeys(jsonMap string, prefix string) (newJsonMap string, err error) {
	dict, err := jsonHelper.ToDict(jsonMap)
	if err != nil {
		return jsonMap, err
	}
	newDict := jsonHelper.Dict{}
	for key, val := range dict {
		if _, exist := newDict[key]; !exist {
			newDict[key] = val
		}
		prefixedKeyParts := strings.SplitN(key, "_", 2)
		if len(prefixedKeyParts) < 2 {
			continue
		}
		keyPrefix, originalKey := prefixedKeyParts[0], prefixedKeyParts[1]
		if keyPrefix != prefix {
			continue
		}
		newDict[originalKey] = val
	}
	newMapBytes, err := json.Marshal(newDict)
	if err != nil {
		return "{}", err
	}
	return string(newMapBytes), nil
}

func (jsonMapUtil *JsonMapUtil) GetFromEnv() (newJsonMap string, err error) {
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

func (jsonMapUtil *JsonMapUtil) ToStringMap(jsonMap string) (newJsonMap string, err error) {
	stringDict, err := jsonHelper.ToStringDict(jsonMap)
	if err != nil {
		return "{}", err
	}
	newMapBytes, err := json.Marshal(stringDict)
	if err != nil {
		return "{}", err
	}
	return string(newMapBytes), nil
}

func (jsonMapUtil *JsonMapUtil) ToVariedStringMap(jsonMap string, keys ...string) (newJsonMap string, err error) {
	variedStringDict, err := jsonHelper.ToStringDict(jsonMap)
	if err != nil {
		return "{}", err
	}
	if len(keys) == 0 {
		for key := range variedStringDict {
			keys = append(keys, key)
		}
	}
	strTransformators := []func(string) string{jsonMapUtil.str.ToKebab, jsonMapUtil.str.ToCamel, jsonMapUtil.str.ToPascal, jsonMapUtil.str.ToSnake, jsonMapUtil.str.ToLower, jsonMapUtil.str.ToUpper, jsonMapUtil.str.ToUpperSnake, jsonMapUtil.str.DoubleQuote, jsonMapUtil.str.SingleQuote}
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

func (jsonMapUtil *JsonMapUtil) ToEnvString(newJsonMap string) (envString string, err error) {
	stringDict, err := jsonHelper.ToStringDict(newJsonMap)
	if err != nil {
		return "", err
	}
	return godotenv.Marshal(stringDict)
}

func (jsonMapUtil *JsonMapUtil) Replace(str string, jsonMapReplacement string) (newStr string, err error) {
	replacementMap, err := jsonHelper.ToStringDict(jsonMapReplacement)
	if err != nil {
		return str, err
	}
	return strutil.StrReplace(str, replacementMap), nil
}
