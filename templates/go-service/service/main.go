package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/config"
	"registry.com/user/serviceName/httpcontroller"
	"registry.com/user/serviceName/rpccontroller"
)

func main() {
	info := config.NewServiceInfo()
	router := gin.Default()
	rpc := communication.NewSimpleRPC(router, info.ServiceURLMap)

	// liveness and readiness
	router.Any("/liveness", httpcontroller.CreateLivenessHandler(info))
	router.Any("/readiness", httpcontroller.CreateReadinessHandler(info))

	// Example: Serve Hello
	router.Any("/", httpcontroller.Hello)
	router.Any("/hello", httpcontroller.Hello)
	router.Any("/hello/:name", httpcontroller.Hello)

	// Example: Register helloRpc
	rpc.RegisterHandler("helloRpc", rpccontroller.Hello)

	// Run RPC Server
	go rpc.Serve()

	// Run HTTP Server
	go router.Run(fmt.Sprintf(":%d", info.HTTPPort))

	// Example: call RPC
	callHelloRPC(rpc, "Kouga")

	forever := make(chan bool)
	<-forever
}

func callHelloRPC(rpc communication.RPC, name string) {
	inputMessage := communication.Message{"name": name}
	log.Printf("[INFO] RPC Call: helloRpc(%v)", inputMessage)
	output, err := rpc.Call("serviceName", "helloRpc", inputMessage)
	if err != nil {
		log.Printf("[ERROR] RPC error: %v", err)
	} else {
		log.Printf("[INFO] RPC Response: %v", output)
	}
}
