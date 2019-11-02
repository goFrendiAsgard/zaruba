package creator

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/state-alchemists/zaruba/config"
)

// Create something from template
func Create(template string, targetPath string, isInteractive bool) error {
	rawTemplatePath, mode := getTemplatePathAndMode(template)
	templatePath := path.Join(config.GetTemplateDir(), rawTemplatePath)
	templateConfig, err := NewTemplateConfig(templatePath)
	if err != nil {
		return err
	}
	modeConfig, err := templateConfig.GetMode(mode)
	if err != nil {
		return err
	}
	log.Printf("%#v", modeConfig)
	os.MkdirAll(targetPath, os.ModePerm)
	return err
}

func getTemplatePathAndMode(template string) (templatePath string, mode string) {
	parts := strings.Split(template, ":")
	if len(parts) == 1 {
		return parts[0], "base"
	}
	return parts[0], strings.Join(parts[1:], ":")
}
