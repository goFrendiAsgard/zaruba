package transport

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

// RmqEventConfig single event config
type RmqEventConfig struct {
	QueueName          string
	ExchangeName       string
	RPCTimeout         int
	DeadLetterExchange string
	DeadLetterQueue    string
	TTL                int
	AutoAck            bool
}

// RmqEventMap map of event config
type RmqEventMap map[string]RmqEventConfig

// GetExchangeName based on event
func (m *RmqEventMap) GetExchangeName(eventName string) (exchangeName string) {
	eventMap := *m
	if config, ok := eventMap[eventName]; ok && config.ExchangeName != "" {
		return config.ExchangeName
	}
	return eventName
}

// GetQueueName based on event
func (m *RmqEventMap) GetQueueName(eventName string) (queueName string) {
	eventMap := *m
	if config, ok := eventMap[eventName]; ok && config.QueueName != "" {
		return config.QueueName
	}
	return eventName
}

// GetDeadLetterExchange based on event
func (m *RmqEventMap) GetDeadLetterExchange(eventName string) (deadLetterExchangeName string) {
	eventMap := *m
	if config, ok := eventMap[eventName]; ok && config.DeadLetterExchange != "" {
		return config.DeadLetterExchange
	}
	return fmt.Sprintf("%s.dlx", m.GetExchangeName(eventName))
}

// GetDeadLetterQueue based on event
func (m *RmqEventMap) GetDeadLetterQueue(eventName string) (deadLetterQueueName string) {
	eventMap := *m
	if config, ok := eventMap[eventName]; ok && config.DeadLetterQueue != "" {
		return config.DeadLetterExchange
	}
	return fmt.Sprintf("%s.dlx", m.GetQueueName(eventName))
}

// GetTTL based on event
func (m *RmqEventMap) GetTTL(eventName string) (TTL int32) {
	eventMap := *m
	if config, ok := eventMap[eventName]; ok && config.TTL > 0 {
		return int32(config.TTL)
	}
	return int32(0)
}

// GetQueueArgs based on event
func (m *RmqEventMap) GetQueueArgs(eventName string) (args amqp.Table) {
	args = make(amqp.Table)
	if m.GetTTL(eventName) <= 0 {
		return nil
	}
	args["x-dead-letter-exchange"] = m.GetDeadLetterExchange(eventName)
	args["x-message-ttl"] = m.GetTTL(eventName)
	return args
}

// GetRPCTimeout based on event
func (m *RmqEventMap) GetRPCTimeout(eventName string) (RPCTimeout time.Duration) {
	eventMap := *m
	if config, ok := eventMap[eventName]; ok && config.RPCTimeout > 0 {
		return time.Duration(config.RPCTimeout)
	}
	return 30000
}

// GetAutoAck based on event
func (m *RmqEventMap) GetAutoAck(eventName string) (autoAck bool) {
	eventMap := *m
	if config, ok := eventMap[eventName]; ok {
		return config.AutoAck
	}
	return false
}
