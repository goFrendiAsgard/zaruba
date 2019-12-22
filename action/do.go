package action

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/dir"
)

// DoActionOption is option for action.Do
type DoActionOption struct {
	MTime       time.Time
	PerformPre  bool
	PerformPost bool
}

// GetDefaultDoOption get new DoOption
func GetDefaultDoOption() DoActionOption {
	return DoActionOption{
		MTime:       time.Time{},
		PerformPre:  true,
		PerformPost: true,
	}
}

// Do with options on projectDir
func Do(actionString, projectDir string, option DoActionOption, arguments ...string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	arguments = append([]string{projectDir}, arguments...)
	// get allDirs
	allDirs, err := dir.GetAllDirs(projectDir)
	if err != nil {
		return
	}
	// pre-action
	if option.PerformPre {
		if err = processAllDirs("pre-"+actionString, allDirs, option, arguments...); err != nil {
			return
		}
	}
	// action
	if err = processAllDirs(actionString, allDirs, option, arguments...); err != nil {
		return
	}
	// post-action
	if option.PerformPost {
		err = processAllDirs("post-"+actionString, allDirs, option, arguments...)
	}
	return
}

func processAllDirs(actionString string, allDirs []string, option DoActionOption, arguments ...string) (err error) {
	// start multiple processDir as go-routines
	errChans := []chan error{}
	for _, dirName := range allDirs {
		errChan := make(chan error)
		go processDir(errChan, actionString, dirName, option, arguments...)
		errChans = append(errChans, errChan)
	}
	// wait all go-routine finished
	for _, errChan := range errChans {
		err = <-errChan
		if err != nil {
			return
		}
	}
	return
}

func processDir(errChan chan error, actionString, dirName string, option DoActionOption, arguments ...string) {
	mTime, err := dir.GetMTime(dirName)
	if err != nil || mTime.Before(option.MTime) {
		errChan <- err
	}
	actionPath := filepath.Join(dirName, fmt.Sprintf("./%s", actionString))
	if _, err := os.Stat(actionPath); err != nil {
		// if file is not exists
		if os.IsNotExist(err) {
			errChan <- nil
			return
		}
		errChan <- err
		return
	}
	err = command.Run(dirName, actionPath, arguments...)
	errChan <- err
}
