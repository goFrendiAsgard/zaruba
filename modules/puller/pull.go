package puller

import (
	"fmt"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/git"
	"github.com/state-alchemists/zaruba/modules/logger"
	"github.com/state-alchemists/zaruba/modules/organizer"
	"github.com/state-alchemists/zaruba/modules/strutil"
)

// Pull monorepo and subtree
func Pull(projectDir string, p *config.ProjectConfig) (err error) {
	// Init project
	if err = git.InitProject(projectDir, p); err != nil {
		return err
	}
	// get current branch
	currentBranch, err := git.GetCurrentBranchName(projectDir)
	if err != nil {
		return err
	}
	// get remotes
	currentGitRemotes, err := git.GetCurrentGitRemotes(projectDir)
	if err != nil {
		return err
	}
	// commit
	git.CommitIfAnyDiff(projectDir, fmt.Sprintf("💀🔽 [PULL] Commit changes at: %s", time.Now().Format(time.RFC3339)))
	logger.Info("🔽 Pulling from main repo")
	if err = command.RunInteractively(projectDir, "git", "pull", "origin", currentBranch); err != nil {
		return err
	}
	// pull from subrepo
	subrepoPrefixMap := p.GetSubrepoPrefixMap(projectDir)
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		if !strutil.IsInArray(componentName, currentGitRemotes) {
			continue
		}
		component, err := p.GetComponentByName(componentName)
		if err != nil {
			return err
		}
		location := component.GetLocation()
		origin := component.GetOrigin()
		branch := currentBranch
		if location == "" || origin == "" {
			continue
		}
		logger.Info("🔽 Pulling from sub-repo %s", componentName)
		if err = command.RunInteractively(projectDir, "git", "subtree", "pull", "--prefix="+subrepoPrefix, "--squash", componentName, branch); err != nil {
			logger.Error("Cannot pull from subrepo `%s`", subrepoPrefix)
		}
	}
	organizer.Organize(projectDir, p)
	return err
}
