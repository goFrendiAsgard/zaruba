package communication

import (
	"log"
)

// RmqRPC rmq rpc
type RmqRPC struct {
	connectionString string
	handlers         map[string]RPCHandler
	logger           *log.Logger
}

// SetLogger set custome logger
func (rpc *RmqRPC) SetLogger(logger *log.Logger) *SimpleRPC {
	rpc.logger = logger
	return rpc
}

// RegisterHandler register servicemap for call
func (rpc *RmqRPC) RegisterHandler(functionName string, handler RPCHandler) {
	rpc.handlers[functionName] = handler
}

// Serve from remote client
func (rpc *RmqRPC) Serve() {

}

// Call remote function
func (rpc *RmqRPC) Call(serviceName, functionName string, input Message) (output Message, err error) {
	// return
	return output, err
}
