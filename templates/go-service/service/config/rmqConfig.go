package config

import "fmt"

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
