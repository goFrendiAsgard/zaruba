package config

import (
	"os"
	"path"
	"path/filepath"
)

// GetTemplateDir retrieve template dir from environment variable
func GetTemplateDir() (absTemplateDir string) {
	templateDir := os.Getenv("ZARUBA_TEMPLATE_DIR")
	if templateDir == "" {
		executable, err := os.Executable()
		if err == nil {
			templateDir = path.Join(filepath.Dir(executable), "templates")
		}
	}
	absTemplateDir, err := filepath.Abs(templateDir)
	if err != nil {
		return templateDir
	}
	return absTemplateDir
}
