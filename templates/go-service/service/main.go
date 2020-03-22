package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/config"
)

func main() {
	info := config.NewServiceInfo()
	router := gin.Default()
	rpc := communication.NewSimpleRPC(router, info.ServiceURLMap)

	registerHTTPHandlers(info, router, rpc)
	registerRPCHandlers(info, router, rpc)

	go rpc.Serve()
	go router.Run(fmt.Sprintf(":%d", info.HTTPPort))

	forever := make(chan bool)
	<-forever
}
