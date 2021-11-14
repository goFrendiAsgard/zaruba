package pathutil

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

type PathUtil struct{}

func NewPathUtil() *PathUtil {
	return &PathUtil{}
}

func (pathUtil *PathUtil) GetRelativePath(basePath, targetPath string) (relativePath string, err error) {
	absBasePath, err := filepath.Abs(basePath)
	if err != nil {
		return "", err
	}
	absTargetPath, err := filepath.Abs(targetPath)
	if err != nil {
		return "", err
	}
	return filepath.Rel(absBasePath, absTargetPath)
}

func (pathUtil *PathUtil) GetDefaultAppName(location string) (appName string, err error) {
	absPath, err := filepath.Abs(location)
	if err != nil {
		return "", err
	}
	baseName := filepath.Base(absPath)
	pattern := regexp.MustCompile(`[^A-Za-z0-9]`)
	spacedBaseName := (pattern.ReplaceAllString(baseName, " "))
	titledBaseName := strings.Title(spacedBaseName)
	appName = strings.ReplaceAll(titledBaseName, " ", "")
	if len(appName) > 0 {
		bts := []byte(appName)
		lc := bytes.ToLower([]byte{bts[0]})
		rest := bts[1:]
		appName = string(bytes.Join([][]byte{lc, rest}, nil))
	}
	return appName, err
}

func (pathUtil *PathUtil) GetEnvFileList(location string) (envFileList []string, err error) {
	absDirPath, err := filepath.Abs(location)
	if err != nil {
		return envFileList, err
	}
	dir, err := os.Open(absDirPath)
	if err != nil {
		return envFileList, err
	}
	defer dir.Close()
	fileList, err := dir.Readdirnames(0)
	if err != nil {
		return envFileList, err
	}
	envFileList = []string{}
	for _, fileName := range fileList {
		if strings.HasSuffix(fileName, ".env") && fileName != ".env" {
			envFileList = append(envFileList, fileName)
		}
	}
	return envFileList, err
}

func (pathUtil *PathUtil) GetEnvByLocation(location string) (envMap map[string]string, err error) {
	absDirPath, err := filepath.Abs(location)
	if err != nil {
		return envMap, err
	}
	files, err := ioutil.ReadDir(absDirPath)
	if err != nil {
		return envMap, err
	}
	envMap = map[string]string{}
	for _, file := range files {
		isDir := file.IsDir()
		if isDir {
			continue
		}
		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".env") && !strings.HasSuffix(fileName, ".env.template") {
			continue
		}
		singleEnvMap, err := godotenv.Read(filepath.Join(absDirPath, fileName))
		if err != nil {
			return envMap, err
		}
		for key, value := range singleEnvMap {
			if _, keyExist := envMap[key]; keyExist {
				continue
			}
			envMap[key] = value
		}
	}
	return envMap, nil
}
