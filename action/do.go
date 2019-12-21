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
	arguments = append([]string{projectDir}, arguments...)
	// get allDirs
	allDirs, err := dir.GetAllDirs(projectDir)
	if err != nil {
		return
	}
	// pre-action
	if err = processAllDirs("pre-"+actionString, allDirs, arguments...); err != nil {
		return
	}
	// action
	if err = processAllDirs(actionString, allDirs, arguments...); err != nil {
		return
	}
	// post-action
	err = processAllDirs("post-"+actionString, allDirs, arguments...)
	return
}

func processAllDirs(actionString string, allDirs []string, arguments ...string) (err error) {
	// start multiple processDir as go-routines
	errChans := []chan error{}
	for _, dir := range allDirs {
		errChan := make(chan error)
		go processDir(errChan, actionString, dir, arguments...)
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

func processDir(errChan chan error, actionString, dir string, arguments ...string) {
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
