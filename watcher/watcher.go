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
func Watch(projectDir string, errChan chan error, stopChan chan bool, arguments ...string) {
	projectDir, err := filepath.Abs(projectDir)
	if err != nil {
		errChan <- err
		return
	}
	log.Printf("[INFO] Watch project `%s` %s", projectDir, stringformat.SprintArgs(arguments))
	organizer.Organize(projectDir, organizer.NewOption().SetMTimeLimitToNow(), arguments...)
	processKillChan := make(chan string) // order to kill process, containing either "shutdown" or "restart"
	executedChan := make(chan bool)      // informing whether execution has been done or not
	terminatedChan := make(chan bool)    // informing whether execution has been terminated or not
	go listen(projectDir, organizer.NewOption().SetMTimeLimitToNow(), executedChan, processKillChan, arguments...)
	go run(projectDir, executedChan, processKillChan, terminatedChan)
	if stop := <-stopChan; stop {
		processKillChan <- "shutdown"
	}
	<-terminatedChan
	errChan <- nil
}

func run(projectDir string, executedChan chan bool, processKillChan chan string, terminatedChan chan bool) {
	errChan := make(chan error)
	stopChan := make(chan bool)
	for {
		go runner.Run(projectDir, stopChan, executedChan, errChan)
		killSignal := <-processKillChan
		stopChan <- true
		if killSignal != "restart" {
			break
		}
	}
	err := <-errChan
	if err != nil {
		log.Printf("[ERROR] Run/terminate process: %s", err)
	}
	terminatedChan <- true
}

func listen(projectDir string, organizerOption *organizer.Option, executedChan chan bool, processKillChan chan string, arguments ...string) {
	// get allDirs
	allDirs, err := getAllDirsTirelessly(projectDir)
	// create watcher, don't give up
	watcher, err := fsnotify.NewWatcher()
	for err != nil {
		log.Printf("[ERROR] Fail to create watcher: %s. Retrying...", err)
		watcher, err = fsnotify.NewWatcher()
	}
	defer watcher.Close()
	// add allDirs to watcher
	addDirsToWatcher(watcher, allDirs)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				continue
			}
			log.Printf("[INFO] Detect event `%s`", event)
			removeDirsFromWatcher(watcher, allDirs)
			organizer.Organize(
				projectDir,
				organizerOption,
				arguments...,
			)
			processKillChan <- "restart"
			<-executedChan
			allDirs, err = getAllDirsTirelessly(projectDir)
			addDirsToWatcher(watcher, allDirs)
			organizerOption = organizer.NewOption().SetMTimeLimitToNow()
		case err, ok := <-watcher.Errors:
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
