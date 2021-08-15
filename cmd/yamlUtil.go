package cmd

func convertYamlObj(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convertYamlObj(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convertYamlObj(v)
		}
	}
	return i
}
