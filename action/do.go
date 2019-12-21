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

// DefaultDoOption is default option used for Do
var DefaultDoOption DoActionOption = DoActionOption{
	MTime:       time.Time{},
	PerformPre:  true,
	PerformPost: true,
}

// Do action on projectDir
func Do(actionString, projectDir string, arguments ...string) (err error) {
	return DoAction(
		actionString,
		projectDir,
		DefaultDoOption,
		arguments...,
	)
}

// DoAction action on projectDir
func DoAction(actionString, projectDir string, option DoActionOption, arguments ...string) (err error) {
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
	for _, dir := range allDirs {
		errChan := make(chan error)
		go processDir(errChan, actionString, dir, option, arguments...)
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

func processDir(errChan chan error, actionString, dir string, option DoActionOption, arguments ...string) {
	actionPath := filepath.Join(dir, fmt.Sprintf("./%s", actionString))
	if _, err := os.Stat(actionPath); err != nil {
		// if file is not exists
		if os.IsNotExist(err) {
			errChan <- nil
			return
		}
		errChan <- err
		return
	}
	err := command.Run(dir, actionPath, arguments...)
	errChan <- err
}
