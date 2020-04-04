package components

import (
	"app/context"
	"app/transport"

	"github.com/gin-gonic/gin"
)

// Setting is data to be injected into SetUp
// TODO: change the implementation to fit your need
type Setting struct {
	Ctx         *context.Context
	Router      *gin.Engine
	Publishers  *Publishers
	Subscribers *Subscribers
	RPCServers  *RPCServers
	RPCClients  *RPCClients
}

// Publishers ...
type Publishers struct {
	Main transport.Publisher
}

// Subscribers ...
type Subscribers struct {
	Main transport.Subscriber
}

// RPCServers ...
type RPCServers struct {
	Main      transport.RPCServer
	Secondary transport.RPCServer
}

// RPCClients ...
type RPCClients struct {
	MainLoopBack      transport.RPCClient
	SecondaryLoopBack transport.RPCClient
}
