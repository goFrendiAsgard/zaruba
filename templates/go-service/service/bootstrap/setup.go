package bootstrap

import "registry.com/user/servicename/components/monitoring"

// Setup everything
func Setup(s *Setting) {
	monitoring.RegisterHealthController(s.Ctx, s.Router)

	// TODO: Add your custom handlers here

}
