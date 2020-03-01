package git

import (
	"log"

	"github.com/state-alchemists/zaruba/modules/command"
)

// Init init git
func Init(projectDir string) (err error) {
	log.Printf("[INFO] Git Init on `%s`", projectDir)
	return command.RunAndRedirect(projectDir, "git", "init")
}
