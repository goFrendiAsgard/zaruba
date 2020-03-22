package main

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/communication"
	"registry.com/user/servicename/rpchandlers"
	"registry.com/user/servicename/servicedesc"
)

func registerRPCHandlers(context *servicedesc.Context, router *gin.Engine, rpc communication.RPC, pubSub communication.PubSub) {
	// EXAMPLE: Register helloRpc handler
	rpc.RegisterHandler("helloRpc", rpchandlers.Hello)
}
