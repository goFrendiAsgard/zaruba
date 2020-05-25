package defaultcomponent

import (
	"app/config"
	"app/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSetup factory to create SetupComponent
func CreateSetup(config *config.Config, router *gin.Engine) core.SetupComponent {
	return func() {
		serviceName := config.ServiceName

		router.Any("/", func(c *gin.Context) {
			// send response
			c.JSON(http.StatusOK, gin.H{
				"service_name": serviceName,
			})
		})

	}
}
