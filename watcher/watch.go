package watcher

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/dir"
	"github.com/state-alchemists/zaruba/hook"
)

// Watch over the project to maintain peace and order
func Watch(project string, stop chan bool) error {
	project, err := filepath.Abs(project)
	if err != nil {
		log.Println(err)
		return err
	}
	go listen(project)
	// wait until stopped
	<-stop
	return err
}

func listen(project string) {
	log.Println("Zaruba load watcher")
	// create watcher, don't give up
	watcher, err := fsnotify.NewWatcher()
	for err != nil {
		watcher, err = fsnotify.NewWatcher()
	}
	// prepare watcher, don't give up
	hc, shell, environ, err := prepareWatcher(watcher, project)
	for err != nil {
		log.Println(err)
		log.Println("Zaruba reload watcher")
		hc, shell, environ, err = prepareWatcher(watcher, project)
	}
	defer watcher.Close()
	// listen
	log.Println("Zaruba is watching")
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				continue
			}
			currentPath := event.Name
			log.Println("Zaruba detect event: ", event)
			// look for matching singleHookConfig
			for watchedPath := range hc {
				// trigger action
				if strings.HasPrefix(currentPath, watchedPath) {
					watcher.Remove(watchedPath)
					err := hc.RunAction(shell, environ, watchedPath)
					if err != nil {
						log.Println(err)
					}
					log.Println("Zaruba reload watcher")
					hc, shell, environ, err = prepareWatcher(watcher, project)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				continue
			}
			log.Println("error:", err)
			continue
		}
	}
}

func prepareWatcher(watcher *fsnotify.Watcher, project string) (hc hook.Config, shell []string, environ []string, err error) {
	// get allDirPaths
	allDirPaths, err := dir.GetAllDirPaths(project)
	if err != nil {
		return
	}
	// get ignoredPaths
	ignoredPaths := getIgnoredPaths(allDirPaths)
	// add event listener to watcher
	err = addDirsToWatcher(watcher, allDirPaths, ignoredPaths)
	if err != nil {
		return
	}
	// create hookConfig
	hc, err = hook.NewCascadedConfig(allDirPaths)
	if err != nil {
		return
	}
	shell = config.GetShell()
	environ = os.Environ()
	return
}

func getIgnoredPaths(allDirPaths []string) []string {
	ignoredPaths := []string{}
	for _, dirPath := range allDirPaths {
		ignoreFile := path.Join(dirPath, config.IgnoreFile)
		ignoreFileContent, err := ioutil.ReadFile(ignoreFile)
		if err != nil {
			continue
		}
		rawIgnorePaths := strings.Split(string(ignoreFileContent), "\n")
		for _, rawIgnorePath := range rawIgnorePaths {
			ignoredPaths = append(ignoredPaths, path.Join(dirPath, rawIgnorePath))
		}
	}
	return ignoredPaths
}

func addDirsToWatcher(watcher *fsnotify.Watcher, allDirPaths []string, ignoredPaths []string) error {
	for _, dirPath := range allDirPaths {
		for _, ignoredPath := range ignoredPaths {
			watcher.Remove(ignoredPath)
		}
		if shouldIgnoreDirPath(dirPath, ignoredPaths) {
			continue
		}
		watcher.Add(dirPath)
	}
	return nil
}

func shouldIgnoreDirPath(dirPath string, ignoredPaths []string) bool {
	for _, ignoredPath := range ignoredPaths {
		if strings.HasPrefix(dirPath, ignoredPath) {
			return true
		}
	}
	return false
}
