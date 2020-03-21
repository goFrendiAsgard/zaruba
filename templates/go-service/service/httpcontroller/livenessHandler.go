package httpcontroller

import (
	"github.com/gin-gonic/gin"
	"registry.com/user/serviceName/config"
)

// CreateLivenessHandler is a factory to create http liveness handler
func CreateLivenessHandler(info *config.ServiceInfo) (handler func(c *gin.Context)) {
	return func(c *gin.Context) {
		c.JSON(info.Status.GetLivenessHTTPCode(), gin.H{
			"serviceName": info.ServiceName,
			"isAlive":     info.Status.IsAlive,
		})
	}
}
