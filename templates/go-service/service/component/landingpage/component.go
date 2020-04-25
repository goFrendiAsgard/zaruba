package landingpage

import (
	"app/config"
	"app/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSetup factory to create SetupComponent
func CreateSetup(app core.App, config *config.Config) core.SetupComponent {
	return func() {
		serviceName := config.ServiceName
		r := app.Router()

		r.Any("/", func(c *gin.Context) {
			// send response
			c.JSON(http.StatusOK, gin.H{
				"service_name": serviceName,
			})
		})

	}
}
