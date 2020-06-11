package defaultcomponent

import (
	"app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Component component definition
type Component struct {
	config *config.Config
	router *gin.Engine
}

// CreateComponent create new component
func CreateComponent(config *config.Config, router *gin.Engine) *Component {
	return &Component{
		config: config,
		router: router,
	}
}

// Setup component
func (comp *Component) Setup() {
	serviceName := comp.config.ServiceName

	comp.router.Any("/", func(c *gin.Context) {
		// send response
		c.JSON(http.StatusOK, gin.H{
			"service_name": serviceName,
		})
	})
}
