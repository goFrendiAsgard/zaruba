package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/communication"
	"registry.com/user/servicename/context"
)

func main() {
	context := context.NewContext()
	config := context.Config
	router := gin.Default()
	pubSub := communication.NewRmqPubSub(config.DefaultRmq.CreateConnectionString()).SetLogger(config.Logger)
	rpc := communication.NewRmqRPC(config.DefaultRmq.CreateConnectionString()).SetLogger(config.Logger)
	// rpc := communication.NewSimpleRPC(router, config.ServiceURLMap).SetLogger(config.Logger)

	registerHTTPHandlers(context, router, rpc, pubSub)
	registerRPCHandlers(context, router, rpc, pubSub)
	registerPubSubHandlers(context, router, rpc, pubSub)

	go pubSub.Start()
	go rpc.Serve()
	go router.Run(fmt.Sprintf(":%d", config.HTTPPort))

	forever := make(chan bool)
	<-forever
}
