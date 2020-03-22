package main

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/communication"
	"registry.com/user/servicename/pubsubhandlers"
	"registry.com/user/servicename/servicedesc"
)

func registerPubSubHandlers(context *servicedesc.Context, router *gin.Engine, rpc communication.RPC, pubSub communication.PubSub) {
	// EXAMPLE: Register helloPubSub handler
	pubSub.RegisterHandler(context.DefaultRmqEvent, pubsubhandlers.CreateHelloHandler(context))
}
