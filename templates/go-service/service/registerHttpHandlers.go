package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/communication"
	"registry.com/user/serviceName/httphandlers"
	"registry.com/user/serviceName/servicedesc"
)

func registerHTTPHandlers(context *servicedesc.Context, router *gin.Engine, rpc communication.RPC, pubSub communication.PubSub) {
	// liveness and readiness handlers
	router.GET("/liveness", httphandlers.CreateLivenessHandler(context))
	router.GET("/readiness", httphandlers.CreateReadinessHandler(context))

	// Default route
	router.Any("/", func(c *gin.Context) { c.String(http.StatusOK, "serviceName") })

	// EXAMPLE: hello
	router.GET("/hello", httphandlers.Hello)
	router.POST("/hello", httphandlers.Hello)
	router.GET("/hello/:name", httphandlers.Hello)
	router.POST("/hello/:name", httphandlers.Hello)

	// EXAMPLE: hello-rpc
	helloRPCHandler := httphandlers.CreateHelloRPCHandler(rpc)
	router.GET("/hello-rpc", helloRPCHandler)
	router.POST("/hello-rpc", helloRPCHandler)
	router.GET("/hello-rpc/:name", helloRPCHandler)
	router.POST("/hello-rpc/:name", helloRPCHandler)

	// EXAMPLE: hello-pubsub
	helloPubSubHandler := httphandlers.CreateHelloPubSubHandler(context, pubSub)
	router.GET("/hello-pubsub", helloPubSubHandler)
	router.POST("/hello-pubsub", helloPubSubHandler)
	router.GET("/hello-pubsub/:name", helloPubSubHandler)
	router.POST("/hello-pubsub/:name", helloPubSubHandler)

	// EXAMPLE: hello-all
	router.GET("/hello-all", httphandlers.CreateHelloAllHandler(context))
}
