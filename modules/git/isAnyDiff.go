package git

import (
	"strings"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// IsAnyDiff get whether there is diff or not
func IsAnyDiff(projectDir string) bool {
	logger.Info("Get current git branch")
	output, err := command.Run(projectDir, "git", "diff", "HEAD", "--exit-code")
	if err != nil {
		return true
	}
	diff := strings.Trim(output, "\r\n ")
	return diff == ""
}
