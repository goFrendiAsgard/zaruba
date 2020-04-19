package core

import (
	"app/transport"
	"log"

	"github.com/gin-gonic/gin"
)

// SetupComponent function to setup component
type SetupComponent func()

// App interface
type App interface {
	Logger() *log.Logger
	Router() *gin.Engine
	GlobalPublisher() transport.Publisher
	LocalPublisher() transport.Publisher
	GlobalSubscriber() transport.Subscriber
	LocalSubscriber() transport.Subscriber
	GlobalRPCServer() transport.RPCServer
	LocalRPCServer() transport.RPCServer
	GlobalRPCClient() transport.RPCClient
	LocalRPCClient() transport.RPCClient
	Liveness() bool
	Readiness() bool
	SetLiveness(liveness bool)
	SetReadiness(readiness bool)
	Setup(componentSetups []SetupComponent)
	Run()
}
