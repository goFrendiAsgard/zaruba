package monitoring

import (
	"app/config"
	"app/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Component component definition
type Component struct {
	config *config.Config
	app    core.App
	router *gin.Engine
}

// CreateComponent create new component
func CreateComponent(config *config.Config, app core.App, router *gin.Engine) *Component {
	return &Component{
		config: config,
		app:    app,
		router: router,
	}
}

// Setup component
func (comp *Component) Setup() {
	serviceName := comp.config.ServiceName

	comp.router.GET("/liveness", func(c *gin.Context) {
		liveness := comp.app.Liveness()
		// send response
		c.JSON(comp.getHTTPCodeByStatus(liveness), gin.H{
			"service_name": serviceName,
			"is_alive":     liveness,
		})
	})

	comp.router.GET("/readiness", func(c *gin.Context) {
		readiness := comp.app.Readiness()
		// send response
		c.JSON(comp.getHTTPCodeByStatus(readiness), gin.H{
			"service_name": serviceName,
			"is_ready":     readiness,
		})
	})

}

func (comp *Component) getHTTPCodeByStatus(status bool) (untyped int) {
	if status {
		return http.StatusOK
	}
	return http.StatusInternalServerError
}
