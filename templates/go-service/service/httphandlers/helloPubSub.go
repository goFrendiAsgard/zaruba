package httphandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/communication"
	"registry.com/user/servicename/servicedesc"
)

// CreateHelloPubSubHandler create helloPubSub handler
func CreateHelloPubSubHandler(context *servicedesc.Context, pubSub communication.PubSub) (handler func(c *gin.Context)) {
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

		// publish to DefaultRmqEvent of current service
		err := pubSub.Publish(context.DefaultRmqEvent, communication.Message{"name": name})
		if err != nil {
			context.Logger.Println(err)
			c.String(http.StatusInternalServerError, "Sending error")
			return
		}
		c.String(http.StatusOK, "Message sent")

	}
}
