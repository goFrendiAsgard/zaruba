package action

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/dir"
)

// Do action on projectDir
func Do(actionString, projectDir string, arguments ...string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	// get allDirs
	allDirs, err := dir.GetAllDirs(projectDir)
	if err != nil {
		return
	}
	// pre-action
	if err = processAllDirs(allDirs, "pre-"+actionString, arguments...); err != nil {
		return
	}
	// action
	if err = processAllDirs(allDirs, actionString, arguments...); err != nil {
		return
	}
	// post-action
	err = processAllDirs(allDirs, "post-"+actionString, arguments...)
	return
}

func processAllDirs(allDirs []string, actionString string, arguments ...string) (err error) {
	// start multiple processDir as go-routines
	errChans := []chan error{}
	for _, dir := range allDirs {
		errChan := make(chan error)
		go processDir(errChan, dir, actionString, arguments...)
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

func processDir(errChan chan error, dir, actionString string, arguments ...string) {
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
