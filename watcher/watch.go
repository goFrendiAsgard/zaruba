package watcher

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

// Watch over the project to maintain peace and order
func Watch(project string, stop chan bool) error {
	// define watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
		return err
	}
	defer watcher.Close()
	// add listener
	log.Println("Zaruba watch for changes")
	go maintain(watcher)
	// add files to watch
	log.Println("Zaruba add path")
	err = watcher.Add(".")
	err = watcher.Add("cmd")
	if err != nil {
		log.Println(err)
	}
	// wait until stopped
	<-stop
	return err
}

func maintain(watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Println("event:", event)
			// detect remove
			if event.Op&fsnotify.Remove == fsnotify.Remove {
				log.Println("removed file:", event.Name)
			}
			// detect write
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
