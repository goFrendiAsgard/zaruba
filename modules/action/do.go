package action

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/file"
)

// Do with options on projectDir
func Do(actionString string, option *Option, arguments ...string) (err error) {
	if err = prepareOption(option); err != nil {
		return err
	}
	// get ignores
	rootWorkDir := option.GetWorkDir()
	ignores := []string{}
	p, loadErr := config.LoadProjectConfig(rootWorkDir)
	if loadErr == nil {
		ignores = p.Ignores
	}
	// get allWorkDirs
	allWorkDirs := []string{rootWorkDir}
	if option.GetIsRecursiveWorkDir() {
		allWorkDirs, err = file.GetAllFiles(option.GetWorkDir(), file.NewOption().SetIsOnlyDir(true).SetIgnores(ignores))
		if err != nil {
			return err
		}
	}
	// pre-action
	if option.GetIsPerformPre() {
		if err = processAllDirs("pre-"+actionString, allWorkDirs, option, arguments...); err != nil {
			return err
		}
	}
	// action
	if option.GetIsPerformAction() {
		if err = processAllDirs(actionString, allWorkDirs, option, arguments...); err != nil {
			return err
		}
	}
	// post-action
	if option.GetIsPerformPost() {
		err = processAllDirs("post-"+actionString, allWorkDirs, option, arguments...)
	}
	return err
}

func prepareOption(option *Option) (err error) {
	// set workDir
	workDir, err := filepath.Abs(option.GetWorkDir())
	if err != nil {
		return err
	}
	option.SetWorkDir(workDir)
	// set scriptDir
	if option.GetScriptDir() == "" {
		option.SetScriptDir(option.GetWorkDir())
	} else {
		var scriptDir string
		scriptDir, err = filepath.Abs(option.GetScriptDir())
		if err != nil {
			return err
		}
		option.SetScriptDir(scriptDir)
	}
	return err
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
			return err
		}
	}
	return err
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
	err = command.RunAndRedirect(workDir, actionPath, arguments...)
	errChan <- err
}

func getActionPath(actionString, workDir string, option *Option) (actionPath string) {
	if option.GetWorkDir() == option.GetScriptDir() {
		return filepath.Join(workDir, fmt.Sprintf("./%s.zaruba", actionString))
	}
	return filepath.Join(option.GetScriptDir(), fmt.Sprintf("./%s.zaruba", actionString))
}
