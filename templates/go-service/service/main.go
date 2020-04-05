package main

import (
	"github.com/gin-gonic/gin"

	"app/bootstrap"
	"app/components"
	"app/context"
	"app/transport"
)

func main() {

	ctx := context.NewContext()
	logger := ctx.Config.Logger
	rmqConnectionString := ctx.Config.RmqConnectionString
	router := gin.Default()

	// define setting
	s := &components.Setting{
		Ctx:    ctx,
		Router: router,
	}
	// main publisher
	s.Publishers.Main = transport.NewRmqPublisher(rmqConnectionString).SetLogger(logger)
	// main subscriber
	s.Subscribers.Main = transport.NewRmqSubscriber(rmqConnectionString).SetLogger(logger)
	// RPC Servers
	s.RPCServers.Main = transport.NewRmqRPCServer(rmqConnectionString).SetLogger(logger)
	s.RPCServers.Secondary = transport.NewSimpleRPCServer(router).SetLogger(logger)
	// RPC Clients
	s.RPCClients.MainLoopBack = transport.NewRmqRPCClient(rmqConnectionString).SetLogger(logger)
	s.RPCClients.SecondaryLoopBack = transport.NewSimpleRPCClient(ctx.Config.LocalServiceAddress).SetLogger(logger)

	bootstrap.Setup(s)
	bootstrap.Run(s)
}
