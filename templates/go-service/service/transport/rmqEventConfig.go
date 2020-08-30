package transport

// RmqEventConfig single event config
type RmqEventConfig struct {
	QueueName    string
	ExchangeName string
}

// RmqEventMap map of event config
type RmqEventMap map[string]RmqEventConfig

// GetExchangeName based on event
func (m *RmqEventMap) GetExchangeName(eventName string) (exchangeName string) {
	eventMap := *m
	if config, ok := eventMap[eventName]; ok {
		return config.ExchangeName
	}
	return eventName
}

// GetQueueName based on event
func (m *RmqEventMap) GetQueueName(eventName string) (queueName string) {
	eventMap := *m
	if config, ok := eventMap[eventName]; ok {
		return config.QueueName
	}
	return eventName
}
