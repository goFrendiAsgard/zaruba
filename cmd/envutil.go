package cmd

import "fmt"

func envCascadePrefix(envMap map[string]string, prefix string) (cascadedEnvMap map[string]string) {
	cascadedEnvMap = map[string]string{}
	for key, val := range envMap {
		prefixedKey := fmt.Sprintf("%s_%s", prefix, key)
		if prefixedVal, exist := envMap[prefixedKey]; exist {
			cascadedEnvMap[key] = prefixedVal
			continue
		}
		cascadedEnvMap[key] = val
	}
	return cascadedEnvMap
}
