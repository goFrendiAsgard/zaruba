package monitoring

import (
	"app/config"
	"app/core"

	"github.com/gin-gonic/gin"
)

// CreateSetup factory to create SetupComponent
func CreateSetup(app core.App, config *config.Config) core.SetupComponent {
	return func() {
		serviceName := config.ServiceName
		r := app.Router()

		r.GET("/liveness", func(c *gin.Context) {
			liveness := app.Liveness()
			// send response
			c.JSON(getHTTPCodeByStatus(liveness), gin.H{
				"service_name": serviceName,
				"is_alive":     liveness,
			})
		})

		r.GET("/readiness", func(c *gin.Context) {
			readiness := app.Readiness()
			// send response
			c.JSON(getHTTPCodeByStatus(readiness), gin.H{
				"service_name": serviceName,
				"is_ready":     readiness,
			})
		})
	}
}
