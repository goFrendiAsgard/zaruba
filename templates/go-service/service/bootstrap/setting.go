package bootstrap

import (
	"app/context"
	"app/transport"

	"github.com/gin-gonic/gin"
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
