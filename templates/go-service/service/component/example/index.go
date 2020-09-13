package example

import (
	"app/config"
	"app/transport"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Component implementation
type Component struct {
	config     *config.Config
	router     *gin.Engine
	publisher  transport.Publisher
	subscriber transport.Subscriber
	rpcServer  transport.RPCServer
	rpcClient  transport.RPCClient
	names      []string
}

// CreateComponent create component
func CreateComponent(config *config.Config, router *gin.Engine, publisher transport.Publisher, subscriber transport.Subscriber, rpcServer transport.RPCServer, rpcClient transport.RPCClient) *Component {
	return &Component{
		router:     router,
		publisher:  publisher,
		subscriber: subscriber,
		rpcServer:  rpcServer,
		rpcClient:  rpcClient,
		config:     config,
		names:      []string{},
	}
}

// Setup component
func (comp *Component) Setup() {
	comp.route()
	comp.registerRPCHandler()
	comp.registerMessageHandler()
}

func (comp *Component) route() {
	// Use the same HTTP Handler for multiple URLS
	comp.router.GET("/hello", comp.handleHTTPHello)
	comp.router.GET("/hello/:name", comp.handleHTTPHello)
	comp.router.POST("/hello", comp.handleHTTPHello)
	// Use HTTP Handler that take state from component
	comp.router.GET("/hello-all", comp.handleHTTPHelloAll)
	// Trigger RPC Call
	comp.router.GET("/hello-rpc", comp.handleHTTPHelloRPC)
	comp.router.GET("/hello-rpc/:name", comp.handleHTTPHelloRPC)
	comp.router.POST("/hello-rpc", comp.handleHTTPHelloRPC)
	// Trigger Publisher Call
	comp.router.GET("/hello-pub", comp.handleHTTPHelloPub)
	comp.router.GET("/hello-pub/:name", comp.handleHTTPHelloPub)
	comp.router.POST("/hello-pub", comp.handleHTTPHelloPub)
}

func (comp *Component) registerRPCHandler() {
	comp.rpcServer.RegisterHandler("helloRPC", comp.handleRPCHello)
}

func (comp *Component) registerMessageHandler() {
	comp.subscriber.RegisterHandler("hello", comp.handleEventHello)
}

func (comp *Component) handleHTTPHello(c *gin.Context) {
	name := getName(c)
	c.String(http.StatusOK, Greet(name))
}

func (comp *Component) handleHTTPHelloAll(c *gin.Context) {
	c.String(http.StatusOK, GreetEveryone(comp.names))
}

func (comp *Component) handleHTTPHelloRPC(c *gin.Context) {
	name := getName(c)
	greetingInterface, err := comp.rpcClient.Call("helloRPC", name)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	greeting, ok := greetingInterface.(string)
	if !ok {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Cannot convert %#v to string", greetingInterface))
		return
	}
	c.String(http.StatusOK, greeting)
}

func (comp *Component) handleHTTPHelloPub(c *gin.Context) {
	name := getName(c)
	err := comp.publisher.Publish("hello", transport.Message{"name": name})
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	c.String(http.StatusOK, "Message sent")
}

func (comp *Component) handleRPCHello(inputs ...interface{}) (greeting interface{}, err error) {
	if len(inputs) == 0 {
		return greeting, errors.New("Message accepted but input is invalid")
	}
	name, success := inputs[0].(string)
	if !success {
		errorMessage := fmt.Sprintf("Cannot convert %#v to string", inputs[0])
		return greeting, errors.New(errorMessage)
	}
	return Greet(name), err
}

func (comp *Component) handleEventHello(msg transport.Message) (err error) {
	name, err := msg.GetString("name")
	if err != nil {
		return err
	}
	comp.names = append(comp.names, name)
	return err
}
