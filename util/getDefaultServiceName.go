package util

import (
	"bytes"
	"path/filepath"
	"regexp"
	"strings"
)

func GetDefaultServiceName(location string) (serviceName string, err error) {
	absPath, err := filepath.Abs(location)
	if err != nil {
		return "", err
	}
	baseName := filepath.Base(absPath)
	pattern := regexp.MustCompile(`[^A-Za-z0-9]`)
	spacedBaseName := (pattern.ReplaceAllString(baseName, " "))
	titledBaseName := strings.Title(spacedBaseName)
	serviceName = strings.ReplaceAll(titledBaseName, " ", "")
	if len(serviceName) > 0 {
		bts := []byte(serviceName)
		lc := bytes.ToLower([]byte{bts[0]})
		rest := bts[1:]
		serviceName = string(bytes.Join([][]byte{lc, rest}, nil))
	}
	return serviceName, err
}
