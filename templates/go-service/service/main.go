package main

import (
	"github.com/gin-gonic/gin"

	"app/bootstrap"
	"app/components"
	"app/context"
	"app/example"
	"app/transport"
)

func main() {

	ctx := context.NewContext()
	logger := ctx.Config.Logger
	rmqConnectionString := ctx.Config.RmqConnectionString
	router := gin.Default()
	s := &components.Setting{
		Ctx:    ctx,
		Router: router,
		Publishers: &components.Publishers{
			Main: transport.NewRmqPublisher(rmqConnectionString).SetLogger(logger),
		},
		Subscribers: &components.Subscribers{
			Main: transport.NewRmqSubscriber(rmqConnectionString).SetLogger(logger),
		},
		RPCServers: &components.RPCServers{
			Main:      transport.NewRmqRPCServer(rmqConnectionString).SetLogger(logger),
			Secondary: transport.NewSimpleRPCServer(router).SetLogger(logger),
		},
		RPCClients: &components.RPCClients{
			MainLoopBack:      transport.NewRmqRPCClient(rmqConnectionString).SetLogger(logger),
			SecondaryLoopBack: transport.NewSimpleRPCClient(ctx.Config.LocalServiceAddress).SetLogger(logger),
		},
	}

	// TODO: remove this.
	example.Setup(s)

	bootstrap.Setup(s)
	bootstrap.Run(s)
}
