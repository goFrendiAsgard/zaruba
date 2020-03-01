package git

import (
	"log"

	"github.com/state-alchemists/zaruba/modules/command"
)

// SubtreeAdd git subtree add
func SubtreeAdd(projectDir, prefix, origin, branch string) (err error) {
	log.Printf("[INFO] Git subtree add `%s` with prefix `%s` and branch `%s`", origin, prefix, branch)
	return command.RunAndRedirect(projectDir, "git", "subtree", "add", "--prefix="+prefix, origin, branch)
}
