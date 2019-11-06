package watcher

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/otiai10/copy"
	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
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
	allDirPaths, err := getAllDirPaths(project)
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
	hookConfig, err := createHookConfig(allDirPaths)
	if err != nil {
		log.Println(err)
	}
	// add listener
	log.Println("Zaruba watch for changes")
	shell := config.GetShell()
	go maintain(watcher, shell, project, &hookConfig)
	// wait until stopped
	<-stop
	return err
}

func createHookConfig(allDirPaths []string) (hookConfig HookConfig, err error) {
	hookConfig = make(HookConfig)
	for _, dirPath := range allDirPaths {
		currentHookConfig, err := NewHookConfig(dirPath)
		if err != nil {
			continue
		}
		for key, val := range currentHookConfig {
			if _, ok := hookConfig[key]; !ok {
				hookConfig[key] = SingleHookConfig{}
			}
			singleConfig := hookConfig[key]
			singleConfig.Dir = val.Dir
			for _, preTrigger := range val.PreTriggers {
				singleConfig.PreTriggers = append(singleConfig.PreTriggers, preTrigger)
			}
			for _, postTrigger := range val.PostTriggers {
				singleConfig.PostTriggers = append(singleConfig.PostTriggers, postTrigger)
			}
			for _, link := range val.Links {
				singleConfig.Links = append(singleConfig.Links, link)
			}
			hookConfig[key] = singleConfig
		}
	}
	return
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

func getAllDirPaths(dirPath string) (allDirPaths []string, err error) {
	allDirPaths = []string{}
	allDirPaths = append(allDirPaths, dirPath)
	subdirs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}
	for _, subdir := range subdirs {
		if !subdir.IsDir() {
			continue
		}
		subdirPath := path.Join(dirPath, subdir.Name())
		subdirPaths, err := getAllDirPaths(subdirPath)
		if err != nil {
			return allDirPaths, err
		}
		allDirPaths = append(allDirPaths, subdirPaths...)
	}
	return
}

func maintain(watcher *fsnotify.Watcher, shell []string, project string, hookConfig *HookConfig) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			currentPath := event.Name
			log.Println("Zaruba detect event: ", event)
			// look for matching singleHookConfig
			for watchedPath, singleHookConfig := range *hookConfig {
				if strings.HasPrefix(currentPath, watchedPath) {
					runSingleHookConfig(shell, watchedPath, singleHookConfig)
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

func runSingleHookConfig(shell []string, hookConfigKey string, singleHookConfig SingleHookConfig) {
	// run pre-triggers
	if err := command.RunMultiple(shell, singleHookConfig.Dir, os.Environ(), singleHookConfig.PreTriggers); err != nil {
		log.Println(err)
		return
	}
	// process links
	for _, link := range singleHookConfig.Links {
		if err := copy.Copy(hookConfigKey, link); err != nil {
			log.Println(err)
			return
		}
	}
	// run post-triggers
	if err := command.RunMultiple(shell, singleHookConfig.Dir, os.Environ(), singleHookConfig.PostTriggers); err != nil {
		log.Println(err)
		return
	}
}
