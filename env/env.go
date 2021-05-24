package env

import (
	"os"
	"strings"
)

func GetEnvFileList(location string) (envFileList []string, err error) {
	dir, err := os.Open(location)
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
