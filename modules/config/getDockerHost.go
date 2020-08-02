package config

import (
	"os"
)

// GetDockerHost get docker host
func GetDockerHost() (dockerHost string) {
	dockerHost = os.Getenv("ZARUBA_DOCKER_HOST")
	if dockerHost == "" {
		return "0.0.0.0"
	}
	return dockerHost
}
