package main

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/communication"
	"registry.com/user/servicename/context"
	"registry.com/user/servicename/rpchandlers"
)

func registerRPCHandlers(context *context.Context, router *gin.Engine, rpc communication.RPC, pubSub communication.PubSub) {
	// EXAMPLE: Register helloRpc handler
	rpc.RegisterHandler("helloRpc", rpchandlers.Hello)
}
