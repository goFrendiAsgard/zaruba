package bootstrap

import (
	"app/components"
	"app/components/monitoring"
)

// Setup everything
func Setup(s *components.Setting) {
	monitoring.Setup(s)

	// TODO: Add your custom handlers here

}
