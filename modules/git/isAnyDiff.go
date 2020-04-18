package git

import (
	"strings"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// IsAnyDiff get whether there is diff or not
func IsAnyDiff(projectDir string) bool {
	logger.Info("Get current git branch")
	output, err := command.Run(projectDir, "git", "--no-pager", "diff", "HEAD")
	if err != nil {
		return false
	}
	diff := strings.Trim(output, "\r\n ")
	return diff != ""
}
