package dir

import (
	"os"
	"time"
)

// GetMTime get modified time of file or dir
func GetMTime(fileName string) (mtime time.Time, err error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return
	}
	// if `filename` is an actual file, return it's modtime
	mtime = fileInfo.ModTime()
	if !fileInfo.IsDir() {
		return
	}
	// if `filename` is a directory, then check for it's subdirectories
	subFileNames, err := GetAllFiles(fileName, DefaultGetFilesOption)
	if err != nil {
		return
	}
	for _, subFileName := range subFileNames {
		fileInfo, err = os.Stat(subFileName)
		subMTime := fileInfo.ModTime()
		if subMTime.After(mtime) {
			mtime = subMTime
		}
	}
	return
}
