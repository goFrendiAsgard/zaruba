package cmd

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/creator"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create <template> <target>",
	Short: "Create artefact",
	Long:  `Zaruba will create artefact based on template`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("template and target required")
		}
		template := args[0]
		target := args[1]
		interactively := false
		for _, arg := args[2:] {
			if arg == "interactively" || arg == "interactive" {
				interactively = true
			}
		}
		if err := creator.Create(template, target, interactively); err != nil {
			log.Fatal(err)
		}
	},
}

func watch() {
	// define watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	// add listener
	log.Println("Zaruba watch for changes")
	go doAction(watcher)
	// add files to watch
	log.Println("Zaruba add path")
	err = watcher.Add(".")
	err = watcher.Add("cmd")
	if err != nil {
		log.Fatal(err)
	}
	// wait forever
	<-done
}

func doAction(watcher *fsnotify.Watcher) {
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
