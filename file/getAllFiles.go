package file

import (
	"io/ioutil"
	"path"
)

// GetAllFiles fetch sub directories in files recursively
func GetAllFiles(fileOrDirName string, option *Option) (allFiles []string, err error) {
	allFiles = []string{}
	allFiles = append(allFiles, fileOrDirName)
	subFiles, err := ioutil.ReadDir(fileOrDirName)
	if err != nil {
		return
	}
	option.SetMaxDepth(option.GetMaxDepth() - 1)
	for _, subFile := range subFiles {
		if !subFile.IsDir() {
			if !option.GetIsOnlyDir() {
				allFiles = append(allFiles, path.Join(fileOrDirName, subFile.Name()))
			}
			continue
		}
		subFilePath := path.Join(fileOrDirName, subFile.Name())
		subfileOrDirNames, err := GetAllFiles(subFilePath, option)
		if err != nil {
			return allFiles, err
		}
		allFiles = append(allFiles, subfileOrDirNames...)
	}
	return
}
