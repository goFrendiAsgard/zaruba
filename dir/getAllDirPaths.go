package dir

import (
	"io/ioutil"
	"path"
)

// GetAllDirPaths fetch sub directories recursively
func GetAllDirPaths(dirPath string) (allDirPaths []string, err error) {
	allDirPaths = []string{}
	allDirPaths = append(allDirPaths, dirPath)
	subdirs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}
	for _, subdir := range subdirs {
		if !subdir.IsDir() {
			continue
		}
		subdirPath := path.Join(dirPath, subdir.Name())
		subdirPaths, err := GetAllDirPaths(subdirPath)
		if err != nil {
			return allDirPaths, err
		}
		allDirPaths = append(allDirPaths, subdirPaths...)
	}
	return
}
