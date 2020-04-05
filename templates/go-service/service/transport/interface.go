package transport

import "log"

// RPCHandler function to handle things
type RPCHandler func(inputs ...interface{}) (output interface{}, err error)

// EventHandler function to handle things
type EventHandler func(input Message) (err error)

// RPCClient interface
type RPCClient interface {
	Call(functionName string, inputs ...interface{}) (output interface{}, err error)
	SetLogger(logger *log.Logger) RPCClient
}

// RPCServer interface
type RPCServer interface {
	RegisterHandler(functionName string, handler RPCHandler) RPCServer
	SetLogger(logger *log.Logger) RPCServer
	Serve()
}

// Publisher interface
type Publisher interface {
	Publish(functionName string, msg Message) (err error)
	SetLogger(logger *log.Logger) Publisher
}

// Subscriber interface
type Subscriber interface {
	RegisterHandler(functionName string, handler EventHandler) Subscriber
	SetLogger(logger *log.Logger) Subscriber
	Subscribe()
}
