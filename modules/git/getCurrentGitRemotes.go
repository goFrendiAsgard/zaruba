package git

import (
	"log"
	"strings"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// GetCurrentGitRemotes get remotes of current projectDir
func GetCurrentGitRemotes(projectDir string) (currentGitRemotes []string, err error) {
	log.Println("[INFO] Get current git remotes")
	output, err := command.Run(projectDir, "git", "remote")
	if err != nil {
		return currentGitRemotes, err
	}
	outputList := strings.Split(output, "\n")
	for _, remote := range outputList {
		remote = strings.Trim(remote, "\r\n ")
		if remote != "" {
			logger.Info("* %s", remote)
			currentGitRemotes = append(currentGitRemotes, remote)
		}
	}
	return currentGitRemotes, err
}
