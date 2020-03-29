package bootstrap

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/context"
	"registry.com/user/servicename/transport"
)

// Setting is data to be injected into SetUp
// TODO: change the implementation to fit your need
type Setting struct {
	Ctx        *context.Context
	Router     *gin.Engine
	Publishers struct {
		Main transport.Publisher
	}
	Subscribers struct {
		Main transport.Subscriber
	}
	RPCServers struct {
		Main      transport.RPCServer
		Secondary transport.RPCServer
	}
	RPCClients struct {
		MainLoopBack      transport.RPCClient
		SecondaryLoopBack transport.RPCClient
	}
}

// NewSetting Create New Setting
func NewSetting() *Setting {
	return &Setting{}
}
