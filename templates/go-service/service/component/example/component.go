package example

import (
	"app/config"
	"app/core"
	"app/transport"

	"github.com/gin-gonic/gin"
)

// Component implementation
type Component struct {
	config     *config.Config
	app        core.App
	router     *gin.Engine
	publisher  transport.Publisher
	subscriber transport.Subscriber
	rpcServer  transport.RPCServer
	rpcClient  transport.RPCClient
	controller *controller
}

// CreateComponent create component
func CreateComponent(config *config.Config, app core.App, router *gin.Engine, publisher transport.Publisher, subscriber transport.Subscriber, rpcServer transport.RPCServer, rpcClient transport.RPCClient) *Component {
	return &Component{
		config:     config,
		app:        app,
		router:     router,
		publisher:  publisher,
		subscriber: subscriber,
		rpcServer:  rpcServer,
		rpcClient:  rpcClient,
		controller: createController(publisher, rpcClient),
	}
}

// Setup component
func (comp *Component) Setup() {
	comp.route()
	comp.registerRPCHandler()
	comp.registerMessageHandler()
}

func (comp *Component) route() {
	controller := comp.controller
	// Use the same HTTP Handler for multiple URLS
	comp.router.GET("/hello", controller.handleHTTPHello)
	comp.router.GET("/hello/:name", controller.handleHTTPHello)
	comp.router.POST("/hello", controller.handleHTTPHello)
	// Use HTTP Handler that take state from component
	comp.router.GET("/hello-all", controller.handleHTTPHelloAll)
	// Trigger RPC Call
	comp.router.GET("/hello-rpc", controller.handleHTTPHelloRPC)
	comp.router.GET("/hello-rpc/:name", controller.handleHTTPHelloRPC)
	comp.router.POST("/hello-rpc", controller.handleHTTPHelloRPC)
	// Trigger Publisher Call
	comp.router.GET("/hello-pub", controller.handleHTTPHelloPub)
	comp.router.GET("/hello-pub/:name", controller.handleHTTPHelloPub)
	comp.router.POST("/hello-pub", controller.handleHTTPHelloPub)
}

func (comp *Component) registerRPCHandler() {
	controller := comp.controller
	comp.rpcServer.RegisterHandler("helloRPC", controller.handleRPCHello)
}

func (comp *Component) registerMessageHandler() {
	controller := comp.controller
	comp.subscriber.RegisterHandler("hello", controller.handleEventHello)
}
