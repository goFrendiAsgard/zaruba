package httphandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/context"
)

// CreateReadinessHandler is a factory to create http liveness handler
func CreateReadinessHandler(context *context.Context) (handler func(c *gin.Context)) {
	return func(c *gin.Context) {
		// get http status
		httpCode := http.StatusOK
		if !context.Status.IsReady {
			httpCode = http.StatusInternalServerError
		}
		// send response
		c.JSON(httpCode, gin.H{
			"servicename": context.Config.ServiceName,
			"isReady":     context.Status.IsReady,
		})
	}
}
