package util

import (
	"path/filepath"
	"regexp"
	"strings"
)

func GetServiceName(location string) (serviceName string, err error) {
	absPath, err := filepath.Abs(location)
	if err != nil {
		return "", err
	}
	baseName := filepath.Base(absPath)
	pattern := regexp.MustCompile(`[^A-Za-z0-9]`)
	spacedBaseName := (pattern.ReplaceAllString(baseName, " "))
	titledBaseName := strings.Title(spacedBaseName)
	serviceName = strings.ReplaceAll(titledBaseName, " ", "")
	return serviceName, err
}
