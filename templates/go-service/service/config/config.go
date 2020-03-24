package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"registry.com/user/servicename/env"
)

// NewConfig initiate new config
func NewConfig() (config *Config) {
	HTTPPort, err := strconv.Atoi(env.Getenv("SERVICENAME_HTTP_PORT", "3000"))
	if err != nil {
		log.Fatal(err)
	}
	RmqPort, err := strconv.Atoi(env.Getenv("RMQ_PORT", "5672"))
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		HTTPPort:        HTTPPort,
		ServiceName:     "servicename",
		Logger:          log.New(os.Stdout, "", log.LstdFlags),
		DefaultRmqEvent: env.Getenv("SERVICENAME_EVENT", "servicename"),
		DefaultRmq: &RmqConfig{
			Host:     env.Getenv("RMQ_HOST", "localhost"),
			Port:     RmqPort,
			User:     env.Getenv("RMQ_USER", "root"),
			Password: env.Getenv("RMQ_PASSWORD", "toor"),
			VHost:    env.Getenv("RMQ_VHOST", "/"),
		},
		ServiceURLMap: map[string]string{
			"servicename": env.Getenv("SERVICENAME_URL", fmt.Sprintf("http://localhost:%d", HTTPPort)),
		},
	}
}

// Config is a general service context
type Config struct {
	HTTPPort        int
	ServiceName     string
	Logger          *log.Logger
	DefaultRmq      *RmqConfig
	DefaultRmqEvent string
	ServiceURLMap   map[string]string
}

// RmqConfig is a rabbitmq configuration
type RmqConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	VHost    string
}

// CreateConnectionString create connectionString of rmqConfig
func (r *RmqConfig) CreateConnectionString() (connectionString string) {
	return fmt.Sprintf("amqp://%s:%s@%s:%d%s", r.User, r.Password, r.Host, r.Port, r.VHost)
}
