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
				ExchangeName: "servicename.exchange.helloRPC",
				QueueName:    "servicename.queue.helloRPC",
			},
			"hello": transport.RmqEventConfig{
				ExchangeName: "servicename.exchange.helloEvent",
				QueueName:    "servicename.queue.helloEvent",
			},
		},
	}
}
