package main

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/config"
	"registry.com/user/serviceName/rpccontroller"
)

func registerRPCHandlers(info *config.ServiceInfo, router *gin.Engine, rpc communication.RPC) {
	// EXAMPLE: Register helloRpc handler
	rpc.RegisterHandler("helloRpc", rpccontroller.Hello)
}
