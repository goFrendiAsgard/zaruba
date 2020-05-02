package example

import (
	"app/config"
	"app/core"
	"app/transport"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateComponent create component
func CreateComponent(app core.App, config *config.Config) *Component {
	return &Component{
		names:  []string{},
		app:    app,
		config: config,
	}
}

// Component implementation
type Component struct {
	names  []string
	app    core.App
	config *config.Config
}

// Setup component
func (comp *Component) Setup() {
	r := comp.app.Router()
	rpcServer := comp.app.GlobalRPCServer()
	subscriber := comp.app.GlobalSubscriber()

	// Use the same HTTP Handler for multiple URLS
	r.GET("/hello", comp.handleHTTPHello)
	r.GET("/hello/:name", comp.handleHTTPHello)
	r.POST("/hello", comp.handleHTTPHello)

	// Use HTTP Handler that take state from component
	r.GET("/hello-all", comp.handleHTTPHelloAll)

	// Trigger RPC Call
	r.GET("/hello-rpc", comp.handleHTTPHelloRPC)
	r.GET("/hello-rpc/:name", comp.handleHTTPHelloRPC)
	r.POST("/hello-rpc", comp.handleHTTPHelloRPC)

	// Trigger Publisher Call
	r.GET("/hello-pub", comp.handleHTTPHelloPub)
	r.GET("/hello-pub/:name", comp.handleHTTPHelloPub)
	r.POST("/hello-pub", comp.handleHTTPHelloPub)

	// Serve RPC
	rpcServer.RegisterHandler("servicename.helloRPC", comp.handleRPCHello)

	// Event
	subscriber.RegisterHandler("servicename.helloEvent", comp.handleEventHello)

}

func (comp *Component) handleHTTPHello(c *gin.Context) {
	name := getName(c)
	c.String(http.StatusOK, Greet(name))
}

func (comp *Component) handleHTTPHelloAll(c *gin.Context) {
	c.String(http.StatusOK, GreetEveryone(comp.names))
}

func (comp *Component) handleHTTPHelloRPC(c *gin.Context) {
	rpcClient := comp.app.GlobalRPCClient()
	name := getName(c)
	greetingInterface, err := rpcClient.Call("servicename.helloRPC", name)
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
	publisher := comp.app.GlobalPublisher()
	name := getName(c)
	err := publisher.Publish("servicename.helloEvent", transport.Message{"name": name})
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
