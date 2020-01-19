package action

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/file"
)

// Do with options on projectDir
func Do(actionString string, option *Option, arguments ...string) (err error) {
	prepareOption(option)
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

func prepareOption(option *Option) (err error) {
	// set workDir
	workDir, err := filepath.Abs(option.GetWorkDir())
	if err != nil {
		return
	}
	option.SetWorkDir(workDir)
	// set scriptDir
	if option.GetScriptDir() == "" {
		option.SetScriptDir(option.GetWorkDir())
	} else {
		var scriptDir string
		scriptDir, err = filepath.Abs(option.GetScriptDir())
		if err != nil {
			return
		}
		option.SetScriptDir(scriptDir)
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
	actionPath := getActionPath(actionString, workDir, option)
	if _, err := os.Stat(actionPath); err != nil {
		// if file is not exists
		if os.IsNotExist(err) {
			errChan <- nil
			return
		}
		errChan <- err
		return
	}
	cmd, err := command.GetCmd(workDir, actionPath, arguments...)
	if err != nil {
		errChan <- err
		return
	}
	err = command.Run(cmd)
	errChan <- err
}

func getActionPath(actionString, workDir string, option *Option) (actionPath string) {
	if option.GetWorkDir() == option.GetScriptDir() {
		return filepath.Join(workDir, fmt.Sprintf("./%s.zaruba", actionString))
	}
	return filepath.Join(option.GetScriptDir(), fmt.Sprintf("./%s.zaruba", actionString))
}
