package git

import (
	"log"

	"github.com/state-alchemists/zaruba/modules/command"
)

// LsRemote perform ls remote, return error if failed
func LsRemote(projectDir, gitURL string) (output string, err error) {
	log.Printf("[INFO] Git ls-remote `%s`", gitURL)
	return command.Run(projectDir, "git", "ls-remote", "-h", gitURL)
}
