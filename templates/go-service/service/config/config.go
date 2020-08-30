package config

import (
	"app/transport"
	"encoding/json"
	"log"
)

// Config is a general service context
type Config struct {
	HTTPPort                   int
	ServiceName                string
	DefaultRmqConnectionString string
	RmqEventMap                *transport.RmqEventMap
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
		HTTPPort:                   GetIntFromEnv("SERVICENAME_HTTP_PORT", 3000),
		ServiceName:                "servicename",
		DefaultRmqConnectionString: GetStrFromEnv("DEFAULT_RMQ_CONNECTION_STRING", "amqp://localhost:5672/"),
		RmqEventMap: &transport.RmqEventMap{
			"helloRPC": transport.RmqEventConfig{
				ExchangeName: "servicename.helloRPC",
				QueueName:    "servicename.helloRPC",
			},
			"hello": transport.RmqEventConfig{
				ExchangeName: "servicename.helloEvent",
				QueueName:    "servicename.helloEvent",
			},
		},
	}
}
