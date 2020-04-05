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
	// load project
	_, currentGitRemotes, err := git.GetCurrentBranchAndRemotes(projectDir, p)
	if err != nil {
		return err
	}
	// commit
	git.Commit(projectDir, fmt.Sprintf("Zaruba: Save before pull from sub-repos at %s", time.Now().Format(time.RFC3339)))
	logger.Info("Pull from main repo")
	if err = command.RunAndRedirect(projectDir, "git", "pull", "origin", "HEAD"); err != nil {
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
		branch := component.GetBranch()
		if location == "" || origin == "" || branch == "" {
			continue
		}
		logger.Info("Pulling from sub-repo %s", componentName)
		if err = command.RunAndRedirect(projectDir, "git", "subtree", "pull", "--prefix="+subrepoPrefix, "--squash", componentName, branch); err != nil {
			return err
		}
	}
	organizer.Organize(projectDir, p, organizer.NewOption())
	return err
}
