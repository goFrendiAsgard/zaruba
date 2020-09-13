package example

import (
	"app/transport"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	names     []string
	publisher transport.Publisher
	rpcClient transport.RPCClient
}

func createController(publisher transport.Publisher, rpcClient transport.RPCClient) *controller {
	return &controller{
		names:     []string{},
		publisher: publisher,
		rpcClient: rpcClient,
	}
}

func (controller *controller) getName(c *gin.Context) string {
	name := c.Param("name")
	if name == "" {
		name = c.Query("name")
	}
	if name == "" {
		name = c.PostForm("name")
	}
	return name
}

func (controller *controller) handleHTTPHello(c *gin.Context) {
	name := controller.getName(c)
	c.String(http.StatusOK, Greet(name))
}

func (controller *controller) handleHTTPHelloAll(c *gin.Context) {
	c.String(http.StatusOK, GreetEveryone(controller.names))
}

func (controller *controller) handleHTTPHelloRPC(c *gin.Context) {
	name := controller.getName(c)
	greetingInterface, err := controller.rpcClient.Call("helloRPC", name)
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

func (controller *controller) handleHTTPHelloPub(c *gin.Context) {
	name := controller.getName(c)
	err := controller.publisher.Publish("hello", transport.Message{"name": name})
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}
	c.String(http.StatusOK, "Message sent")
}

func (controller *controller) handleRPCHello(inputs ...interface{}) (greeting interface{}, err error) {
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

func (controller *controller) handleEventHello(msg transport.Message) (err error) {
	name, err := msg.GetString("name")
	if err != nil {
		return err
	}
	controller.names = append(controller.names, name)
	return err
}
