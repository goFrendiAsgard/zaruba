package git

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Checkout checkout to branchName
func Checkout(projectDir, branchName string, newBranch bool) (err error) {
	logger.Info("Git checkout to `%s`", branchName)
	if newBranch {
		return command.RunInteractively(projectDir, "git", "checkout", "-b", branchName)
	}
	return command.RunInteractively(projectDir, "git", "checkout", branchName)
}
