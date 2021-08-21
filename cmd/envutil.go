package cmd

import (
	"strings"
)

func envCascadePrefix(envMap map[string]string, envPrefix string) (cascadedEnvMap map[string]string) {
	cascadedEnvMap = map[string]string{}
	for key, val := range envMap {
		cascadedEnvMap[key] = val
	}
	for prefixedKey, val := range envMap {
		prefixedKeyParts := strings.SplitN(prefixedKey, "_", 2)
		if len(prefixedKeyParts) < 2 {
			continue
		}
		prefix, key := prefixedKeyParts[0], prefixedKeyParts[1]
		if prefix != envPrefix {
			continue
		}
		cascadedEnvMap[key] = val
	}
	return cascadedEnvMap
}
