package servicedesc

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/env"
)

// NewContext initiate new Context
func NewContext() (context *Context) {
	HTTPPort, err := strconv.Atoi(env.Getenv("SERVICENAME_HTTP_PORT", "3000"))
	if err != nil {
		log.Fatal(err)
	}
	RmqPort, err := strconv.Atoi(env.Getenv("RMQ_PORT", "5672"))
	if err != nil {
		log.Fatal(err)
	}
	return &Context{
		HTTPPort:        HTTPPort,
		ServiceName:     "serviceName",
		DefaultRmqEvent: env.Getenv("SERVICENAME_EVENT", "serviceName"),
		DefaultRmq: &RmqConfig{
			Host:     env.Getenv("RMQ_HOST", "localhost"),
			Port:     RmqPort,
			User:     env.Getenv("RMQ_USER", "root"),
			Password: env.Getenv("RMQ_PASSWORD", "toor"),
			VHost:    env.Getenv("RMQ_VHOST", "/"),
		},
		ServiceURLMap: map[string]string{
			"serviceName": env.Getenv("SERVICENAME_URL", fmt.Sprintf("http://localhost:%d", HTTPPort)),
		},
		Status: &Status{
			IsAlive: true,
			IsReady: true,
		},
		Logger:     log.New(os.Stdout, "", log.LstdFlags),
		LocalCache: communication.Message{},
	}
}

// Context is a general service context
type Context struct {
	HTTPPort        int
	ServiceName     string
	DefaultRmq      *RmqConfig
	DefaultRmqEvent string
	ServiceURLMap   map[string]string
	Status          *Status
	Logger          *log.Logger
	LocalCache      communication.Message
}
