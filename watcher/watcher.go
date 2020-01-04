package watcher

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/organizer"
	"github.com/state-alchemists/zaruba/runner"
	"github.com/state-alchemists/zaruba/stringformat"
)

// Watch projectDir
func Watch(projectDir string, stopChan chan bool, errChan chan error, arguments ...string) {
	projectDir, err := filepath.Abs(projectDir)
	if err != nil {
		errChan <- err
		return
	}
	log.Printf("[INFO] Watch project `%s` %s", projectDir, stringformat.SprintArgs(arguments))
	organizer.Organize(projectDir, organizer.NewOption().SetMTimeLimitToNow(), arguments...)
	go listen(projectDir, organizer.NewOption().SetMTimeLimitToNow(), arguments...)
	<-stopChan
	errChan <- nil
}

func listen(projectDir string, organizerOption *organizer.Option, arguments ...string) {
	// run services and wait until executed
	runnerStopChan := make(chan bool)
	runnerErrChan := make(chan error)
	runnerExecutedChan := make(chan bool)
	go runner.Run(projectDir, runnerStopChan, runnerExecutedChan, runnerErrChan)
	<-runnerExecutedChan
	// define `isListening`
	isListening := true
	// get allDirs
	allDirs, err := getAllDirsTirelessly(projectDir)
	// create watcher, don't give up
	w, err := fsnotify.NewWatcher()
	for err != nil {
		log.Printf("[ERROR] Fail to create watcher: %s. Retrying...", err)
		w, err = fsnotify.NewWatcher()
	}
	defer w.Close()
	// add allDirs to watcher
	addDirsToWatcher(w, allDirs)
	for {
		select {
		case event, ok := <-w.Events:
			if !ok || !isListening {
				continue
			}
			log.Printf("[INFO] Detect event `%s`", event)
			removeDirsFromWatcher(w, allDirs)
			// stop services
			runnerStopChan <- true
			<-runnerErrChan
			// re run services
			go runner.Run(projectDir, runnerStopChan, runnerExecutedChan, runnerErrChan)
			<-runnerExecutedChan
			// re-organize
			organizerOption = organizer.NewOption().SetMTimeLimitToNow()
			organizer.Organize(
				projectDir,
				organizerOption,
				arguments...,
			)
			allDirs, err = getAllDirsTirelessly(projectDir)
			addDirsToWatcher(w, allDirs)
		case err, ok := <-w.Errors:
			if !ok {
				continue
			}
			log.Printf("[ERROR] Watcher error: %s. Continue to listen...", err)
		}
	}
}

func getAllDirsTirelessly(projectDir string) (allDirs []string, err error) {
	allDirs, err = file.GetAllFiles(projectDir, file.NewOption().SetIsOnlyDir(true))
	for err != nil {
		log.Printf("[ERROR] Fail to get list of directories: %s. Retrying...", err)
		allDirs, err = file.GetAllFiles(projectDir, file.NewOption().SetIsOnlyDir(true))
	}
	return
}

func removeDirsFromWatcher(watcher *fsnotify.Watcher, allDirs []string) {
	for _, dirPath := range allDirs {
		watcher.Remove(dirPath)
	}
}

func addDirsToWatcher(watcher *fsnotify.Watcher, allDirs []string) {
	for _, dirPath := range allDirs {
		watcher.Add(dirPath)
	}
}
