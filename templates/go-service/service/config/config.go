package config

import (
	"fmt"
	"log"
	"os"
)

// NewConfig initiate new config
func NewConfig() (config *Config) {
	servicePort := GetIntFromEnv("SERVICENAME_HTTP_PORT", 3000)
	rmqConnectionString := GetRmqConnectionString(
		GetStrFromEnv("RMQ_HOST", "localhost"),
		GetIntFromEnv("RMQ_PORT", 5672),
		GetStrFromEnv("RMQ_USER", "root"),
		GetStrFromEnv("RMQ_PASSWORD", "toor"),
		GetStrFromEnv("RMQ_VHOST", "/"),
	)
	return &Config{
		HTTPPort:            servicePort,
		ServiceName:         "servicename",
		Logger:              log.New(os.Stdout, "", log.LstdFlags),
		RmqConnectionString: rmqConnectionString,
		LocalServiceAddress: fmt.Sprintf("http://localhost:%d", servicePort),
	}
}

// Config is a general service context
type Config struct {
	HTTPPort            int
	ServiceName         string
	Logger              *log.Logger
	RmqConnectionString string
	LocalServiceAddress string
}
