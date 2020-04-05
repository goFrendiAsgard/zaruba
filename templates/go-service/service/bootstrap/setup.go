package bootstrap

import (
	"app/components"
	"app/components/example"
	"app/components/monitoring"
)

// Setup application
func Setup(s *components.Setting) {
	monitoring.Setup(s)
	example.Setup(s)
}
