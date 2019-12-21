package dir

import (
	"io/ioutil"
	"path"
)

// GetFilesOption is configuration used for second parameter of GetAllFiles
type GetFilesOption struct {
	MaxDepth int
	OnlyDir  bool
}

// DefaultGetFilesOption is predefined option
var DefaultGetFilesOption GetFilesOption = GetFilesOption{
	MaxDepth: 100,
	OnlyDir:  false,
}

// GetAllFiles fetch sub directories in files recursively
func GetAllFiles(dirName string, option GetFilesOption) (allFiles []string, err error) {
	allFiles = []string{}
	allFiles = append(allFiles, dirName)
	subFiles, err := ioutil.ReadDir(dirName)
	if err != nil {
		return
	}
	option.MaxDepth = option.MaxDepth - 1
	for _, subFile := range subFiles {
		if !subFile.IsDir() {
			if !option.OnlyDir {
				allFiles = append(allFiles, path.Join(dirName, subFile.Name()))
			}
			continue
		}
		subFilePath := path.Join(dirName, subFile.Name())
		subdirNames, err := GetAllFiles(subFilePath, option)
		if err != nil {
			return allFiles, err
		}
		allFiles = append(allFiles, subdirNames...)
	}
	return
}
