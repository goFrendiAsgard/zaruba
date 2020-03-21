package config

import (
	"fmt"
	"log"
	"strconv"
)

// NewServiceInfo initiate new ServiceInfo
func NewServiceInfo() (info *ServiceInfo) {
	HTTPPort, err := strconv.Atoi(Getenv("SERVICENAME_HTTP_PORT", "3000"))
	if err != nil {
		log.Fatal(err)
	}
	RmqPort, err := strconv.Atoi(Getenv("RMQ_PORT", "5672"))
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceInfo{
		HTTPPort:        HTTPPort,
		ServiceName:     "serviceName",
		DefaultRmqEvent: Getenv("SERVICENAME_EVENT", "serviceName"),
		DefaultRmq: &RmqConfig{
			Host:     Getenv("RMQ_HOST", "localhost"),
			Port:     RmqPort,
			User:     Getenv("RMQ_USER", "root"),
			Password: Getenv("RMQ_PASSWORD", "toor"),
			VHost:    Getenv("RMQ_VHOST", "/"),
		},
		Status: &ServiceStatus{
			IsAlive: true,
			IsReady: true,
		},
		ServiceURLMap: map[string]string{
			"serviceName": Getenv("SERVICENAME_URL", fmt.Sprintf("http://localhost:%d", HTTPPort)),
		},
	}
}

// ServiceInfo is a general service information
type ServiceInfo struct {
	HTTPPort        int
	ServiceName     string
	DefaultRmq      *RmqConfig
	DefaultRmqEvent string
	Status          *ServiceStatus
	ServiceURLMap   map[string]string
}
