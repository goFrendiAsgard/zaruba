package watcher

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/state-alchemists/zaruba/modules/file"
	"github.com/state-alchemists/zaruba/modules/organizer"
	"github.com/state-alchemists/zaruba/modules/runner"
	"github.com/state-alchemists/zaruba/modules/strutil"
)

// Watch projectDir
func Watch(projectDir string, stopChan chan bool, errChan chan error, arguments ...string) {
	projectDir, err := filepath.Abs(projectDir)
	if err != nil {
		errChan <- err
		return
	}
	log.Printf("[INFO] Watch project `%s` %s", projectDir, strutil.SprintArgs(arguments))
	organizer.Organize(projectDir, organizer.NewOption().SetMTimeLimitToNow(), arguments...)
	// start to listen for changes and do appropriate actions
	listenerStopChan := make(chan bool)
	listenerErrChan := make(chan error)
	go listen(projectDir, listenerStopChan, listenerErrChan, arguments...)
	// wait for stop request
	<-stopChan
	// trigger listener to stop
	listenerStopChan <- true
	// wait for error from listener and return it
	err = <-listenerErrChan
	errChan <- err
}

func listen(projectDir string, listenerStopChan chan bool, listenerErrChan chan error, arguments ...string) {
	// run services and wait until executed
	runnerStopChan := make(chan bool)
	runnerErrChan := make(chan error)
	runnerExecutedChan := make(chan bool)
	go runner.Run(projectDir, runnerStopChan, runnerExecutedChan, runnerErrChan)
	<-runnerExecutedChan
	isListening := true
	allDirs := getAllDirsTirelessly(projectDir)
	w := getNewWatcherTirelessly()
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
			organizer.Organize(projectDir, organizer.NewOption().SetMTimeLimitToNow(), arguments...)
			allDirs = getAllDirsTirelessly(projectDir)
			addDirsToWatcher(w, allDirs)
		case err, ok := <-w.Errors:
			if !ok {
				continue
			}
			log.Printf("[ERROR] Watcher error: %s. Continue to listen...", err)
		case stop, ok := <-listenerStopChan:
			if !ok {
				continue
			}
			runnerStopChan <- stop
			err := <-runnerErrChan
			listenerErrChan <- err
			break
		}
	}
}

func getNewWatcherTirelessly() (w *fsnotify.Watcher) {
	w, err := fsnotify.NewWatcher()
	for err != nil {
		log.Printf("[ERROR] Fail to create watcher: %s. Retrying...", err)
		w, err = fsnotify.NewWatcher()
	}
	return w
}

func getAllDirsTirelessly(projectDir string) (allDirs []string) {
	allDirs, err := file.GetAllFiles(projectDir, file.NewOption().SetIsOnlyDir(true))
	for err != nil {
		log.Printf("[ERROR] Fail to get list of directories: %s. Retrying...", err)
		allDirs, err = file.GetAllFiles(projectDir, file.NewOption().SetIsOnlyDir(true))
	}
	return allDirs
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
