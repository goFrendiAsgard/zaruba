package communication

// NewRmqPubSub create new RmqPubSub
func NewRmqPubSub(connectionString string) *RmqPubSub {
	return &RmqPubSub{
		connectionString: connectionString,
		handlers:         map[string]PubSubHandler{},
	}
}

// RmqPubSub for publish and subscribe
type RmqPubSub struct {
	connectionString string
	handlers         map[string]PubSubHandler
}

// RegisterHandler register servicemap for call
func (pubSub *SimpleRPC) RegisterHandler(functionName string, handler PubSubHandler) {
	pubSub.handlers[functionName] = handler
}
