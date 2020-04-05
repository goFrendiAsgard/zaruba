package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// GetStrFromEnv getting env
func GetStrFromEnv(env, defaultValue string) (value string) {
	if value, ok := os.LookupEnv(env); ok {
		return value
	}
	return defaultValue
}

// GetIntFromEnv getting env
func GetIntFromEnv(env string, defaultValue int) (value int) {
	if strVal, ok := os.LookupEnv(env); ok {
		value, err := strconv.Atoi(strVal)
		if err != nil {
			log.Fatal(err)
		}
		return value
	}
	return defaultValue
}

// GetRmqConnectionString get RMQConnectionString
func GetRmqConnectionString(host string, port int, user, password, vHost string) (connectionString string) {
	return fmt.Sprintf("amqp://%s:%s@%s:%d%s", user, password, host, port, vHost)
}
