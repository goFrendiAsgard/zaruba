package communication

// Message for RPC and pub-sub
type Message map[string]interface{}

// RPCHandler function to handle RPC
type RPCHandler func(input Message) (output Message, err error)

// RPC interface
type RPC interface {
	Call(serviceName, functionName string, input Message) (output Message, err error)
	Serve()
	RegisterHandler(functionName string, handler RPCHandler)
}

// PubSubHandler function to handle RPC
type PubSubHandler func(input Message) (err error)

// PubSub interface
type PubSub interface {
	Publish(eventName string, message Message)
	Subscribe()
	RegisterHandler(eventName string, handler PubSubHandler)
}
