package pathutil

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/state-alchemists/zaruba/jsonutil"
)

type PathUtil struct {
	json *jsonutil.JsonUtil
}

func NewPathUtil(jsonUtil *jsonutil.JsonUtil) *PathUtil {
	return &PathUtil{
		json: jsonUtil,
	}
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

func (pathUtil *PathUtil) GetEnvFileList(location string) (jsonList string, err error) {
	fileList, err := PathGetEnvFileList(location)
	if err != nil {
		return "[]", err
	}
	return pathUtil.json.FromStringList(fileList)
}

func (pathUtil *PathUtil) GetPortConfigByLocation(location string) (jsonList string, err error) {
	envMap, err := PathGetEnvByLocation(location)
	if err != nil {
		return "[]", err
	}
	ports := []string{}
	for key, val := range envMap {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		// NOTES: well known ports are excluded
		// more about IANA's valid ports: https://datatracker.ietf.org/doc/html/rfc6335#section-6
		if intVal >= 1024 && intVal <= 65535 {
			ports = append(ports, fmt.Sprintf("{{ .GetEnv \"%s\" }}", key))
		}
	}
	return pathUtil.json.FromStringList(ports)
}

func (pathUtil *PathUtil) GetEnvByLocation(location string) (jsonMap string, err error) {
	envMap, err := PathGetEnvByLocation(location)
	if err != nil {
		return "{}", err
	}
	return pathUtil.json.FromStringDict(envMap)
}
