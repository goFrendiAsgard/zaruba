package watcher

import (
	"log"
	"os"
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
	// define watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
		return err
	}
	defer watcher.Close()
	allDirPaths, err := dir.GetAllDirPaths(project)
	if err != nil {
		log.Println(err)
	}
	// add files to watch
	log.Println("Zaruba add paths to watcher")
	err = addDirToWatcher(watcher, allDirPaths)
	if err != nil {
		log.Println(err)
	}
	// create hookConfig
	log.Println("Zaruba load configs")
	hc, err := hook.NewCascadedConfig(allDirPaths)
	if err != nil {
		log.Println(err)
	}
	// add listener
	log.Println("Zaruba watch for changes")
	shell := config.GetShell()
	environ := os.Environ()
	go maintain(watcher, shell, environ, project, &hc)
	// wait until stopped
	<-stop
	return err
}

func addDirToWatcher(watcher *fsnotify.Watcher, allDirPaths []string) error {
	for _, dirPath := range allDirPaths {
		err := watcher.Add(dirPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func maintain(watcher *fsnotify.Watcher, shell []string, environ []string, project string, hc *hook.Config) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			currentPath := event.Name
			log.Println("Zaruba detect event: ", event)
			// look for matching singleHookConfig
			for watchedPath := range *hc {
				if strings.HasPrefix(currentPath, watchedPath) {
					err := hc.RunAction(shell, environ, watchedPath)
					if err != nil {
						log.Println(err)
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
