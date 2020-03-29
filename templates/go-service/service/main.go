package main

import (
	"github.com/gin-gonic/gin"

	"registry.com/user/servicename/bootstrap"
	"registry.com/user/servicename/context"
	"registry.com/user/servicename/example"
	"registry.com/user/servicename/transport"
)

func main() {

	s := bootstrap.NewSetting()

	s.Ctx = context.NewContext()
	s.Router = gin.Default()

	logger := s.Ctx.Config.Logger
	rmqConnectionString := s.Ctx.Config.RmqConnectionString

	s.Publishers.Main = transport.NewRmqPublisher(rmqConnectionString).SetLogger(logger)
	s.Subscribers.Main = transport.NewRmqSubscriber(rmqConnectionString).SetLogger(logger)

	s.RPCServers.Main = transport.NewRmqRPCServer(rmqConnectionString).SetLogger(logger)
	s.RPCClients.MainLoopBack = transport.NewRmqRPCClient(rmqConnectionString).SetLogger(logger)

	s.RPCServers.Secondary = transport.NewSimpleRPCServer(s.Router).SetLogger(logger)
	s.RPCClients.SecondaryLoopBack = transport.NewSimpleRPCClient(s.Ctx.Config.LocalServiceAddress).SetLogger(logger)

	// TODO: remove the example, and implement your own
	example.SetUp(s)

	bootstrap.Setup(s)
	bootstrap.Run(s)
}
