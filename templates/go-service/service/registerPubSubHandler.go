package main

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/pubsubhandlers"
	"registry.com/user/serviceName/servicedesc"
)

func registerPubSubHandlers(context *servicedesc.Context, router *gin.Engine, rpc communication.RPC, pubSub communication.PubSub) {
	// EXAMPLE: Register helloPubSub handler
	pubSub.RegisterHandler(context.DefaultRmqEvent, pubsubhandlers.CreateHelloHandler(context))
}
