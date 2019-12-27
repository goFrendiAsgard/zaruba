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
func Do(actionString, scriptDir string, option *Option, arguments ...string) (err error) {
	// make scriptDir absolute
	scriptDir, err = filepath.Abs(scriptDir)
	if err != nil {
		return
	}
	// if option.workDir is empty, set it to scriptDir
	option, err = option.SetScriptDir(scriptDir)
	if err != nil {
		return
	}
	arguments = append([]string{option.GetWorkDir()}, arguments...)
	// get allWorkDirs
	allWorkDirs := []string{option.GetWorkDir()}
	if option.GetIsRecursiveWorkDir() {
		allWorkDirs, err = file.GetAllFiles(option.GetWorkDir(), file.NewOption().SetIsOnlyDir(true))
		if err != nil {
			return
		}
	}
	// pre-action
	if option.GetIsPerformPre() {
		if err = processAllDirs("pre-"+actionString, allWorkDirs, option, arguments...); err != nil {
			return
		}
	}
	// action
	if option.GetIsPerformAction() {
		if err = processAllDirs(actionString, allWorkDirs, option, arguments...); err != nil {
			return
		}
	}
	// post-action
	if option.GetIsPerformPost() {
		err = processAllDirs("post-"+actionString, allWorkDirs, option, arguments...)
	}
	return
}

func processAllDirs(actionString string, allWorkDirs []string, option *Option, arguments ...string) (err error) {
	// start multiple processDir as go-routines
	errChans := []chan error{}
	for _, workDir := range allWorkDirs {
		errChan := make(chan error)
		go processDir(errChan, actionString, workDir, option, arguments...)
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

func processDir(errChan chan error, actionString, workDir string, option *Option, arguments ...string) {
	mTime, err := file.GetMTime(workDir)
	if err != nil || mTime.Before(option.mTimeLimit) {
		errChan <- err
	}
	actionPath := filepath.Join(workDir, fmt.Sprintf("./%s.zaruba", actionString))
	if _, err := os.Stat(actionPath); err != nil {
		// if file is not exists
		if os.IsNotExist(err) {
			errChan <- nil
			return
		}
		errChan <- err
		return
	}
	log.Printf("[INFO] Invoke `%s` on `%s` %s", actionString, workDir, format.SprintArgs(arguments))
	err = command.Run(workDir, actionPath, arguments...)
	errChan <- err
}
