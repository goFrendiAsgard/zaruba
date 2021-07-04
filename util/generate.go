package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Generate(templateLocation, destinationLocation string, replacementMap map[string]string) (err error) {
	absTemplateLocation, err := filepath.Abs(templateLocation)
	if err != nil {
		return err
	}
	absDestinationLocation, err := filepath.Abs(destinationLocation)
	if err != nil {
		return err
	}
	return filepath.Walk(absTemplateLocation,
		func(templatePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativePath := templatePath[len(absTemplateLocation):]
			destinationPath := filepath.Join(absDestinationLocation, replaceByMap(relativePath, replacementMap))
			fileMode := info.Mode()
			if info.IsDir() {
				os.Mkdir(destinationPath, fileMode)
				return nil
			}
			contentB, err := ioutil.ReadFile(templatePath)
			if err != nil {
				return err
			}
			content := string(contentB)
			newContent := replaceByMap(content, replacementMap)
			if newContent == content {
				return nil
			}
			return ioutil.WriteFile(destinationPath, []byte(newContent), fileMode)
		},
	)
}

func replaceByMap(s string, replacementMap map[string]string) (newS string) {
	newS = s
	for key, val := range replacementMap {
		newS = strings.ReplaceAll(newS, key, val)
	}
	return newS
}
