package git

import (
	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Init init git
func Init(projectDir string) (err error) {
	logger.Info("Git Init on `%s`", projectDir)
	return command.RunInteractively(projectDir, "git", "init")
}
