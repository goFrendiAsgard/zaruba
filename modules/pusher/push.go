package pusher

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

// Push monorepo and subtree
func Push(projectDir string, p *config.ProjectConfig) (err error) {
	// init project
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
	// organize
	if err = organizer.Organize(projectDir, p); err != nil {
		return err
	}
	// commit
	git.Commit(projectDir, fmt.Sprintf("Zaruba: Save before push to sub-repos at %s", time.Now().Format(time.RFC3339)))
	logger.Info("Push to main repo")
	if err = command.RunAndRedirect(projectDir, "git", "push", "origin", currentBranch); err != nil {
		return err
	}
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
		if location == "" || origin == "" || branch == "" {
			continue
		}
		logger.Info("Push to sub-repo %s", componentName)
		if err = command.RunAndRedirect(projectDir, "git", "subtree", "push", "--prefix="+subrepoPrefix, componentName, branch); err != nil {
			return err
		}
	}
	return err
}
