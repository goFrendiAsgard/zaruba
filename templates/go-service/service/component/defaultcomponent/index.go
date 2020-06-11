package defaultcomponent

import (
	"app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Component component definition
type Component struct {
	Config *config.Config
	Router *gin.Engine
}

// CreateComponent create new component
func CreateComponent(config *config.Config, router *gin.Engine) *Component {
	return &Component{
		Config: config,
		Router: router,
	}
}

// Setup component
func (comp *Component) Setup() {
	serviceName := comp.Config.ServiceName

	comp.Router.Any("/", func(c *gin.Context) {
		// send response
		c.JSON(http.StatusOK, gin.H{
			"service_name": serviceName,
		})
	})
}
