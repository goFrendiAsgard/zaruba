package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/servicedesc"
)

func main() {
	context := servicedesc.NewContext()
	router := gin.Default()
	rpc := communication.NewSimpleRPC(router, context.ServiceURLMap).SetLogger(context.Logger)
	pubSub := communication.NewRmqPubSub(context.DefaultRmq.CreateConnectionString()).SetLogger(context.Logger)

	registerHTTPHandlers(context, router, rpc, pubSub)
	registerRPCHandlers(context, router, rpc, pubSub)
	registerPubSubHandlers(context, router, rpc, pubSub)

	go pubSub.Start()
	go rpc.Serve()
	go router.Run(fmt.Sprintf(":%d", context.HTTPPort))

	forever := make(chan bool)
	<-forever
}
