package watcher

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/organizer"
)

// Watch projectDir
func Watch(projectDir string, errChan chan error, stopChan chan bool, arguments ...string) {
	projectDir, err := filepath.Abs(projectDir)
	if err != nil {
		errChan <- err
		return
	}
	organizer.Organize(projectDir, organizer.NewOption().SetMTimeLimitToNow(), arguments...)
	go listen(projectDir, organizer.NewOption().SetMTimeLimitToNow(), arguments...)
	if stop := <-stopChan; stop {
		errChan <- nil
	}
}

func listen(projectDir string, organizerOption *organizer.Option, arguments ...string) {
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
			log.Printf("[INFO] Detect event: %s", event)
			removeDirsFromWatcher(watcher, allDirs)
			organizer.Organize(
				projectDir,
				organizerOption,
				arguments...,
			)
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
	allDirs, err = file.GetAllFiles(projectDir, file.NewOption().SetOnlyDir(true))
	for err != nil {
		log.Printf("[ERROR] Fail to get list of directories: %s. Retrying...", err)
		allDirs, err = file.GetAllFiles(projectDir, file.NewOption().SetOnlyDir(true))
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
