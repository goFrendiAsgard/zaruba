package main

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/communication"
	"registry.com/user/servicename/context"
	"registry.com/user/servicename/pubsubhandlers"
)

func registerPubSubHandlers(context *context.Context, router *gin.Engine, rpc communication.RPC, pubSub communication.PubSub) {
	// EXAMPLE: Register helloPubSub handler
	pubSub.RegisterHandler(context.Config.DefaultRmqEvent, pubsubhandlers.CreateHelloHandler(context))
}
