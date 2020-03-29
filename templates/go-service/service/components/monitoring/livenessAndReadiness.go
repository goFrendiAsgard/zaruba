package monitoring

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"registry.com/user/servicename/context"
)

// RegisterHealthController registerHealthController
func RegisterHealthController(ctx *context.Context, router *gin.Engine) {

	router.GET("/liveness", func(c *gin.Context) {
		// get http status
		httpCode := http.StatusOK
		if !ctx.Status.IsAlive {
			httpCode = http.StatusInternalServerError
		}
		// send response
		c.JSON(httpCode, gin.H{
			"service_name": ctx.Config.ServiceName,
			"is_alive":     ctx.Status.IsAlive,
		})
	})

	router.GET("/readiness", func(c *gin.Context) {
		// get http status
		httpCode := http.StatusOK
		if !ctx.Status.IsReady {
			httpCode = http.StatusInternalServerError
		}
		// send response
		c.JSON(httpCode, gin.H{
			"service_name": ctx.Config.ServiceName,
			"is_alive":     ctx.Status.IsReady,
		})
	})

}
