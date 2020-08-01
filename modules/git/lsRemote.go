package git

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// LsRemote perform ls remote, return error if failed
func LsRemote(projectDir, gitURL string) (output string, err error) {
	logger.Info("Git ls-remote `%s`", gitURL)
	return command.RunAndReturn(projectDir, "git", "ls-remote", "-h", gitURL)
}
