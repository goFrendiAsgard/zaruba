package git

import (
	"strings"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// GetCurrentBranchName get current branch name
func GetCurrentBranchName(projectDir string) (branchName string, err error) {
	logger.Info("Get current git branch")
	output, err := command.Run(projectDir, "git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return branchName, err
	}
	branchName = strings.Trim(output, "\r\n ")
	return branchName, err
}
