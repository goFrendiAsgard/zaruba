package httphandlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/context"
)

// CreateHelloAllHandler create helloPubSub handler
func CreateHelloAllHandler(context *context.Context) (handler func(c *gin.Context)) {
	return func(c *gin.Context) {
		greeting := "Hello, everyone..."
		context.InitLocalCache("names", []string{})
		names, _ := context.LocalCache.GetStringArray("names")
		if len(names) > 0 {
			greeting = fmt.Sprintf("Hello %s, and everyone...", strings.Join(names, ", "))
		}
		c.String(http.StatusOK, greeting)
	}
}
