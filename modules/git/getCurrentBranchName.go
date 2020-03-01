package git

import (
	"log"
	"strings"

	"github.com/state-alchemists/zaruba/modules/command"
)

// GetCurrentBranchName get current branch name
func GetCurrentBranchName(projectDir string) (branchName string, err error) {
	log.Println("[INFO] Get current branch")
	output, err := command.Run(projectDir, "git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return branchName, err
	}
	branchName = strings.Trim(output, "\r\n ")
	return branchName, err
}
