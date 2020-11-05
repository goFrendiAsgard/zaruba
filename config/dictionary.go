package config

import "strings"

// Dictionary is advance map
type Dictionary map[string]string

// GetSubKeys get subkeys
func (d Dictionary) GetSubKeys(parentKeys ...string) (subKeys []string) {
	return d.GetSubKeysBySeparator("::", parentKeys...)
}

// GetSubKeysBySeparator get subkeys by separator
func (d Dictionary) GetSubKeysBySeparator(separator string, parentKeys ...string) (subKeys []string) {
	seen := map[string]bool{}
	parentKey := strings.Join(parentKeys, separator)
	prefixLength := len(parentKey) + len(separator)
	subKeys = []string{}
	for key := range d {
		if !strings.HasPrefix(key, parentKey) {
			continue
		}
		childKey := key[prefixLength:]
		if childKey == "" {
			continue
		}
		childKeyParts := strings.SplitN(childKey, separator, 2)
		subkey := childKeyParts[0]
		seen[subkey] = true
	}
	for key := range seen {
		subKeys = append(subKeys, key)
	}
	return subKeys
}

// GetValue of dictionary
func (d Dictionary) GetValue(keys ...string) (val string) {
	return d.GetValueBySeparator("::", keys...)
}

// GetValueBySeparator of dictionary
func (d Dictionary) GetValueBySeparator(separator string, keys ...string) (val string) {
	key := strings.Join(keys, separator)
	return d[key]
}
