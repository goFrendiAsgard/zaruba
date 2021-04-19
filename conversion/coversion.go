package conversion

func NormalizeMapStringValue(stringMap map[string]string) (genericMap map[string]interface{}) {
	genericMap = map[string]interface{}{}
	for key, val := range stringMap {
		genericMap[key] = val
	}
	return genericMap
}

func NormalizeMapListStringValue(listStringMap map[string][]string) (genericMap map[string]interface{}) {
	genericMap = map[string]interface{}{}
	for key, val := range listStringMap {
		genericMap[key] = val
	}
	return genericMap
}
