package httpcontroller

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/config"
)

// CreateReadinessHandler is a factory to create http liveness handler
func CreateReadinessHandler(info *config.ServiceInfo) (handler func(c *gin.Context)) {
	return func(c *gin.Context) {
		c.JSON(info.Status.GetReadinessHTTPCode(), gin.H{
			"serviceName": info.ServiceName,
			"isReady":     info.Status.IsReady,
		})
	}
}
