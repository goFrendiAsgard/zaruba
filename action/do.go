package action

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/format"
)

// Do with options on projectDir
func Do(actionString, projectDir string, option *Option, arguments ...string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	arguments = append([]string{projectDir}, arguments...)
	// get allDirs
	allDirs, err := file.GetAllFiles(projectDir, file.NewOption().SetOnlyDir(true))
	if err != nil {
		return
	}
	// pre-action
	if option.GetPerformPre() {
		if err = processAllDirs("pre-"+actionString, allDirs, option, arguments...); err != nil {
			return
		}
	}
	// action
	if option.GetPerformAction() {
		if err = processAllDirs(actionString, allDirs, option, arguments...); err != nil {
			return
		}
	}
	// post-action
	if option.GetPerformPost() {
		err = processAllDirs("post-"+actionString, allDirs, option, arguments...)
	}
	return
}

func processAllDirs(actionString string, allDirs []string, option *Option, arguments ...string) (err error) {
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

func processDir(errChan chan error, actionString, dirName string, option *Option, arguments ...string) {
	mTime, err := file.GetMTime(dirName)
	if err != nil || mTime.Before(option.mTimeLimit) {
		errChan <- err
	}
	actionPath := filepath.Join(dirName, fmt.Sprintf("./%s.zaruba", actionString))
	if _, err := os.Stat(actionPath); err != nil {
		// if file is not exists
		if os.IsNotExist(err) {
			errChan <- nil
			return
		}
		errChan <- err
		return
	}
	log.Printf("[INFO] Invoke `%s` on `%s` %s", actionString, dirName, format.SprintArgs(arguments))
	err = command.Run(dirName, actionPath, arguments...)
	errChan <- err
}
