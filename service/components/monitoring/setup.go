package monitoring

import (
	"app/components"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup registerHealthController
func Setup(s *components.Setting) {

	s.Router.GET("/liveness", func(c *gin.Context) {
		// get http status
		httpCode := http.StatusOK
		if !s.Ctx.Status.IsAlive {
			httpCode = http.StatusInternalServerError
		}
		// send response
		c.JSON(httpCode, gin.H{
			"service_name": s.Ctx.Config.ServiceName,
			"is_alive":     s.Ctx.Status.IsAlive,
		})
	})

	s.Router.GET("/readiness", func(c *gin.Context) {
		// get http status
		httpCode := http.StatusOK
		if !s.Ctx.Status.IsReady {
			httpCode = http.StatusInternalServerError
		}
		// send response
		c.JSON(httpCode, gin.H{
			"service_name": s.Ctx.Config.ServiceName,
			"is_alive":     s.Ctx.Status.IsReady,
		})
	})

}
