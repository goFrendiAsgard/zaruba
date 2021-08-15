package cmd

import (
	"encoding/json"
	"fmt"
)

func convertToMapString(mapInterface map[string]interface{}) (mapString map[string]string) {
	mapString = map[string]string{}
	for key, value := range mapInterface {
		if strValue, ok := value.(string); ok {
			mapString[key] = strValue
			continue
		}
		if valB, err := json.Marshal(value); err == nil {
			mapString[key] = string(valB)
			continue
		}
		mapString[key] = fmt.Sprintf("%v", value)
	}
	return mapString
}
