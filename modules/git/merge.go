package git

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Merge checkout to branchName
func Merge(projectDir, branchName string) (err error) {
	logger.Info("Git merge `%s`", branchName)
	return command.RunAndRedirect(projectDir, "git", "merge", "--squash", branchName)
}
