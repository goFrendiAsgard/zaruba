package util

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/str"
)

func Generate(sourceTemplatePath, destinationPath string, replacementMap map[string]string) (err error) {
	absSourceTemplatePath, err := filepath.Abs(sourceTemplatePath)
	if err != nil {
		return err
	}
	absDestinationPath, err := filepath.Abs(destinationPath)
	if err != nil {
		return err
	}
	return filepath.Walk(absSourceTemplatePath,
		func(absSourceLocation string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativeLocation := absSourceLocation[len(absSourceTemplatePath):]
			absDestinationLocation := filepath.Join(absDestinationPath, str.ReplaceByMap(relativeLocation, replacementMap))
			fileMode := info.Mode()
			if info.IsDir() {
				os.Mkdir(absDestinationLocation, fileMode)
				return nil
			}
			contentB, err := ioutil.ReadFile(absSourceLocation)
			if err != nil {
				return err
			}
			content := string(contentB)
			newContent := str.ReplaceByMap(content, replacementMap)
			if newContent == content {
				_, err = copyFile(absSourceLocation, absDestinationLocation)
				return err
			}
			return ioutil.WriteFile(absDestinationLocation, []byte(newContent), fileMode)
		},
	)
}

func copyFile(src, dst string) (byteCount int64, err error) {
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
