package config

import (
	"encoding/json"
	"log"
)

// Config is a general service context
type Config struct {
	HTTPPort                  int
	ServiceName               string
	GlobalRmqConnectionString string
	LocalRmqConnectionString  string
}

// ToString change config into string
func (c *Config) ToString() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal("[ERROR]", err)
	}
	return string(b)
}

// CreateConfig initiate new config
func CreateConfig() (config *Config) {
	return &Config{
		HTTPPort:                  GetIntFromEnv("SERVICENAME_HTTP_PORT", 3000),
		ServiceName:               "servicename",
		GlobalRmqConnectionString: GetStrFromEnv("GLOBAL_RMQ_CONNECTION_STRING", "amqp://localhost:5672/"),
		LocalRmqConnectionString:  GetStrFromEnv("LOCAL_RMQ_CONNECTION_STRING", "amqp://localhost:5672/"),
	}
}
