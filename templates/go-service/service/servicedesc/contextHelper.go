package servicedesc

import "fmt"

// InitLocalCache set LocalCache
func (context *Context) InitLocalCache(key string, val interface{}) {
	if _, exists := context.LocalCache[key]; !exists {
		context.LocalCache[key] = val
	}
}

// Status represent liveness and readiness of service
type Status struct {
	IsAlive bool
	IsReady bool
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
