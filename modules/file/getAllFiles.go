package file

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// GetAllFiles fetch sub-files and sub-directories recursively
func GetAllFiles(dirName string, option *Option) (allFiles []string, err error) {
	resultChan := make(chan resultOfGetFile)
	go getAllFiles(dirName, option, resultChan)
	result := <-resultChan
	return result.fileNames, result.err
}

type resultOfGetFile struct {
	err       error
	fileNames []string
}

func getAllFiles(dir string, option *Option, resultChan chan resultOfGetFile) {
	fileNames := []string{dir}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		resultChan <- resultOfGetFile{err: err, fileNames: fileNames}
		return
	}
	// recursive
	subResultChanList := []chan resultOfGetFile{}
	for _, file := range files {
		absName := filepath.Join(dir, file.Name())
		shouldExclude := false
		for _, exclude := range option.GetIgnores() {
			if strings.HasPrefix(absName, exclude) || absName == filepath.Join(dir, ".git") {
				shouldExclude = true
			}
		}
		if shouldExclude || (option.GetIsOnlyDir() && !file.IsDir()) {
			continue
		}
		fileNames = append(fileNames, absName)
		// if directory, do recursive
		if !file.IsDir() {
			continue
		}
		subResultChan := make(chan resultOfGetFile)
		subResultChanList = append(subResultChanList, subResultChan)
		go getAllFiles(absName, option, subResultChan)
	}
	// handle result
	for _, subResultChan := range subResultChanList {
		subResult := <-subResultChan
		err = subResult.err
		if err != nil {
			resultChan <- resultOfGetFile{err: err, fileNames: fileNames}
			return
		}
		fileNames = append(fileNames, subResult.fileNames...)
	}
	resultChan <- resultOfGetFile{err: err, fileNames: fileNames}
}
