package httphandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/communication"
)

// CreateHelloRPCHandler create helloRPC handler
func CreateHelloRPCHandler(rpc communication.RPC) (handler func(c *gin.Context)) {
	return func(c *gin.Context) {
		// get name
		name := c.Param("name")
		if name == "" {
			name = c.Query("name")
		}
		if name == "" {
			name = c.PostForm("name")
		}
		if name == "" {
			name = "world"
		}

		// call RPC
		output, err := rpc.Call("servicename", "helloRpc", communication.Message{"name": name})
		if err != nil {
			c.String(http.StatusInternalServerError, "RPC Call error")
			return
		}
		// get greeting from RPC output
		greeting, err := output.GetString("greeting")
		if err != nil {
			c.String(http.StatusInternalServerError, "Convertion error")
			return
		}
		c.String(http.StatusOK, greeting)

	}
}
