package config

import (
	"fmt"
	"os"
	"strings"
)

// VirtualEnv ...
type VirtualEnv struct {
	env           map[string]string
	envParseOrder []string
	nativeEnv     map[string]string
}

// Add add to virtualEnv
func (v *VirtualEnv) Add(key, value string) {
	v.env[key] = v.ParseString(value)
	v.envParseOrder = append(v.envParseOrder, key)
}

// ParseString parse string
func (v *VirtualEnv) ParseString(str string) (newStr string) {
	newStr = str
	// replace with current env
	for _, key := range v.envParseOrder {
		newStr = v.replace(newStr, key, v.env[key])
	}
	// replace with native env (from os)
	for key, value := range v.nativeEnv {
		newStr = v.replace(newStr, key, value)
	}
	return newStr
}

func (v *VirtualEnv) replace(str, key, replacement string) (newStr string) {
	newStr = strings.ReplaceAll(str, fmt.Sprintf("${%s}", key), replacement)
	newStr = strings.ReplaceAll(newStr, fmt.Sprintf("$%s", key), replacement)
	return newStr
}

// GetEnv of virtualEnv
func (v *VirtualEnv) GetEnv() (env map[string]string) {
	env = map[string]string{}
	for key, value := range v.env {
		env[key] = value
	}
	return env
}

// CreateVirtualEnv get new virtualEnv
func CreateVirtualEnv() (v *VirtualEnv) {
	v = &VirtualEnv{
		env:           map[string]string{},
		envParseOrder: []string{},
		nativeEnv:     map[string]string{},
	}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		v.nativeEnv[pair[0]] = pair[1]
	}
	return v
}
