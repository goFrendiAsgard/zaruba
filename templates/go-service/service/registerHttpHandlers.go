package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/config"
	"registry.com/user/serviceName/httpcontroller"
)

func registerHTTPHandlers(info *config.ServiceInfo, router *gin.Engine, rpc communication.RPC) {
	// liveness and readiness handlers
	router.Any("/liveness", httpcontroller.CreateLivenessHandler(info))
	router.Any("/readiness", httpcontroller.CreateReadinessHandler(info))

	// Default route
	router.Any("/", func(c *gin.Context) { c.String(http.StatusOK, "serviceName") })

	// EXAMPLE: hello
	router.Any("/hello", httpcontroller.Hello)
	router.Any("/hello/:name", httpcontroller.Hello)

	// EXAMPLE: hello-rpc
	router.Any("/hello-rpc", httpcontroller.CreateHelloRPCHandler(rpc))
	router.Any("/hello-rpc/:name", httpcontroller.CreateHelloRPCHandler(rpc))
}
