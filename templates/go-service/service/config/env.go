package config

import "os"

// Getenv getting env
func Getenv(env, defaultValue string) (value string) {
	if value, ok := os.LookupEnv(env); ok {
		return value
	}
	return defaultValue
}
