package transport

// RPCHandler function to handle things
type RPCHandler func(inputs ...interface{}) (output interface{}, err error)

// EventHandler function to handle things
type EventHandler func(input Message) (err error)

// RPCClient interface
type RPCClient interface {
	Call(functionName string, inputs ...interface{}) (output interface{}, err error)
}

// RPCServer interface
type RPCServer interface {
	RegisterHandler(functionName string, handler RPCHandler) RPCServer
	Serve(errChan chan error)
}

// Publisher interface
type Publisher interface {
	Publish(functionName string, msg Message) (err error)
}

// Subscriber interface
type Subscriber interface {
	RegisterHandler(functionName string, handler EventHandler) Subscriber
	Subscribe(errChan chan error)
}
