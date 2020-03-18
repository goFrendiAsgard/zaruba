package git

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Commit add and commit
func Commit(projectDir, message string) (err error) {
	logger.Info("Git add and git commit with messsage `%s`", message)
	if err = command.RunAndRedirect(projectDir, "git", "add", ".", "-A"); err != nil {
		return err
	}
	err = command.RunAndRedirect(projectDir, "git", "commit", "-m", message)
	return err
}
