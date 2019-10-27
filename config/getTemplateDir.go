package config

import (
	"os"
	"path"
)

// GetTemplateDir retrieve template dir from environment variable
func GetTemplateDir() string {
	templateDir := os.Getenv("ZARUBA_TEMPLATE_DIR")
	if templateDir == "" {
		executable, err := os.Executable()
		if err == nil {
			templateDir = path.Join(path.Base(executable), "templates")
		}
	}
	return templateDir
}
