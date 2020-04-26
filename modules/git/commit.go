package git

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Commit add and commit
func Commit(projectDir, message string) {
	logger.Info("Git add and git commit with messsage `%s`", message)
	if err := command.RunAndRedirect(projectDir, "git", "add", ".", "-A"); err != nil {
		logger.Error("%s", err)
		return
	}
	if err := command.RunAndRedirect(projectDir, "git", "commit", "-m", message); err != nil {
		logger.Error("%s", err)
	}
}

// CommitIfAnyDiff add and commit
func CommitIfAnyDiff(projectDir, message string) {
	if !IsAnyDiff(projectDir) {
		logger.Info("Nothing to commit")
		return
	}
	Commit(projectDir, message)
}
