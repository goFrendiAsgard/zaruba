package httphandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/context"
)

// CreateLivenessHandler is a factory to create http liveness handler
func CreateLivenessHandler(context *context.Context) (handler func(c *gin.Context)) {
	return func(c *gin.Context) {
		// get http status
		httpCode := http.StatusOK
		if !context.Status.IsAlive {
			httpCode = http.StatusInternalServerError
		}
		// send response
		c.JSON(httpCode, gin.H{
			"servicename": context.Config.ServiceName,
			"isAlive":     context.Status.IsAlive,
		})
	}
}
