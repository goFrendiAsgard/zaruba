package dir

import (
	"io/ioutil"
	"path"
)

// GetAllDirs fetch sub directories recursively
func GetAllDirs(dirPath string) (allDirs []string, err error) {
	allDirs = []string{}
	allDirs = append(allDirs, dirPath)
	subdirs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}
	for _, subdir := range subdirs {
		if !subdir.IsDir() {
			continue
		}
		subdirPath := path.Join(dirPath, subdir.Name())
		subdirPaths, err := GetAllDirs(subdirPath)
		if err != nil {
			return allDirs, err
		}
		allDirs = append(allDirs, subdirPaths...)
	}
	return
}
