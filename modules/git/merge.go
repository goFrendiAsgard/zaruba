package git

import (
	"log"

	"github.com/state-alchemists/zaruba/modules/command"
)

// Merge checkout to branchName
func Merge(projectDir, branchName string) (err error) {
	log.Printf("[INFO] Git merge `%s`", branchName)
	return command.RunAndRedirect(projectDir, "git", "merge", "--squash", branchName)
}
