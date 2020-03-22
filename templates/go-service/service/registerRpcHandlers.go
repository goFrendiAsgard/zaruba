package main

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/rpchandlers"
	"registry.com/user/serviceName/servicedesc"
)

func registerRPCHandlers(context *servicedesc.Context, router *gin.Engine, rpc communication.RPC, pubSub communication.PubSub) {
	// EXAMPLE: Register helloRpc handler
	rpc.RegisterHandler("helloRpc", rpchandlers.Hello)
}
