package bootstrap

import (
	"app/components"
	"app/components/example"
	"app/components/monitoring"
)

// Setup components
func Setup(s *components.Setting) {
	monitoring.Setup(s)
	example.Setup(s)
}
