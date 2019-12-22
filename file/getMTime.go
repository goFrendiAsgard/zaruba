package file

import (
	"os"
	"time"
)

// GetMTime get modified time of file or dir
func GetMTime(fileName string) (mTime time.Time, err error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return
	}
	// if `filename` is an actual file, return it's modtime
	mTime = fileInfo.ModTime()
	if !fileInfo.IsDir() {
		return
	}
	// if `filename` is a directory, then check for it's subdirectories
	subFileNames, err := GetAllFiles(fileName, NewOption())
	if err != nil {
		return
	}
	for _, subFileName := range subFileNames {
		fileInfo, err = os.Stat(subFileName)
		if err != nil {
			return
		}
		subMTime := fileInfo.ModTime()
		if subMTime.After(mTime) {
			mTime = subMTime
		}
	}
	return
}
