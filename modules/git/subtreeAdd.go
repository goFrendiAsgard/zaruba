package git

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// SubtreeAdd git subtree add
func SubtreeAdd(projectDir, prefix, origin, branch string) (err error) {
	logger.Info("Git subtree add `%s` with prefix `%s` and branch `%s`", origin, prefix, branch)
	return command.RunInteractively(projectDir, "git", "subtree", "add", "--prefix="+prefix, origin, branch)
}
