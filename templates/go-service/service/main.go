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
	rmqConnectionString := config.DefaultRmq.CreateConnectionString()
	logger := config.Logger

	router := gin.Default()
	pubSub := communication.NewRmqPubSub(rmqConnectionString).SetLogger(logger)
	rpc := communication.NewRmqRPC(rmqConnectionString).SetLogger(logger)
	// rpc := communication.NewSimpleRPC(router, config.ServiceURLMap).SetLogger(logger)

	registerHTTPHandlers(context, router, rpc, pubSub)
	registerRPCHandlers(context, router, rpc, pubSub)
	registerPubSubHandlers(context, router, rpc, pubSub)

	go pubSub.Start()
	go rpc.Serve()
	go router.Run(fmt.Sprintf(":%d", config.HTTPPort))
	logger.Println(fmt.Sprintf("Ruant at port %d", config.HTTPPort))

	forever := make(chan bool)
	<-forever
}
