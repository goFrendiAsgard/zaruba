package git

import (
	"log"

	"github.com/state-alchemists/zaruba/modules/command"
)

// Checkout checkout to branchName
func Checkout(projectDir, branchName string, newBranch bool) (err error) {
	log.Printf("[INFO] Checkout to %s", branchName)
	if newBranch {
		return command.RunAndRedirect(projectDir, "git", "checkout", "-b", branchName)
	}
	return command.RunAndRedirect(projectDir, "git", "checkout", branchName)
}
