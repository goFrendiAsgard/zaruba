package file

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/str"
)

func Copy(src, dst string) (byteCount int64, err error) {
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

func ReadText(fileName string) (text string, err error) {
	contentB, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(contentB), nil
}

func WriteText(fileName string, text string, fileMode os.FileMode) (err error) {
	return ioutil.WriteFile(fileName, []byte(text), fileMode)
}

func ReadLines(fileName string) (lines []string, err error) {
	content, err := ReadText(fileName)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(content, "\n"), nil
}

func WriteLines(fileName string, lines []string, fileMode os.FileMode) (err error) {
	content := strings.Join(lines, "\n")
	return WriteText(fileName, content, fileMode)
}

func ListDir(dirPath string) (fileNames []string, err error) {
	fileNames = []string{}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fileNames, err
	}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames, nil
}

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
			content, err := ReadText(absSourceLocation)
			if err != nil {
				return err
			}
			newContent := str.ReplaceByMap(content, replacementMap)
			if newContent == content {
				_, err = Copy(absSourceLocation, absDestinationLocation)
				return err
			}
			return WriteText(absDestinationLocation, newContent, fileMode)
		},
	)
}
