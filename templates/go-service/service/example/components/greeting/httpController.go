package greeting

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/context"
	"registry.com/user/servicename/transport"
)

func getName(c *gin.Context) string {
	name := c.Param("name")
	if name == "" {
		name = c.Query("name")
	}
	if name == "" {
		name = c.PostForm("name")
	}
	return name
}

// GreetHTTPController is HTTP controller to handle greeting
func GreetHTTPController(c *gin.Context) {
	name := getName(c)
	c.String(http.StatusOK, Greet(name))
}

// CreateGreetEveryoneHTTPController is factory to create HTTP Controller
func CreateGreetEveryoneHTTPController(ctx *context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx.InitLocalCache("names", []string{})
		names, _ := ctx.LocalCache.GetStringArray("names")
		c.String(http.StatusOK, GreetEveryone(names))
	}
}

// CreateGreetRPCHTTPController is factory to create HTTP Controller
func CreateGreetRPCHTTPController(rpcClient transport.RPCClient, functionName string) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := getName(c)
		greetingInterface, err := rpcClient.Call(functionName, name)
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
}

// CreateGreetPublishHTTPController is factory to create HTTP Controller
func CreateGreetPublishHTTPController(publisher transport.Publisher, eventName string) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := getName(c)
		err := publisher.Publish(eventName, transport.Message{"name": name})
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
			return
		}
		c.String(http.StatusOK, "Message sent")
	}
}
