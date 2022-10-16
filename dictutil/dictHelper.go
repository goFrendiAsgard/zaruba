package dictutil

import (
	"encoding/json"
	"sort"
)

func DictGetSortedKeys(any interface{}) (sortedKeys []string, err error) {
	jsonB, err := json.Marshal(any)
	if err != nil {
		return sortedKeys, err
	}
	var dict map[string]interface{}
	err = json.Unmarshal(jsonB, &dict)
	if err != nil {
		return sortedKeys, err
	}
	sortedKeys = []string{}
	for envKey := range dict {
		sortedKeys = append(sortedKeys, envKey)
	}
	sort.Strings(sortedKeys)
	return sortedKeys, nil
}
