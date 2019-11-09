package initiator

import (
	"log"
	"os"

	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/dir"
	"github.com/state-alchemists/zaruba/hook"
)

// Init everything once
func Init(project string) (err error) {
	allDirPaths, err := dir.GetAllDirPaths(project)
	if err != nil {
		return
	}
	// create hookConfig
	log.Println("Zaruba load configs")
	hc, err := hook.NewCascadedConfig(allDirPaths)
	if err != nil {
		return
	}
	sortedKeys := hc.GetSortedKeys()
	shell := config.GetShell()
	environ := os.Environ()
	for _, key := range sortedKeys {
		if err = hc.RunAction(shell, environ, key); err != nil {
			return
		}
	}
	return
}
