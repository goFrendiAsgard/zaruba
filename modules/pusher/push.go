package pusher

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/git"
	"github.com/state-alchemists/zaruba/modules/organizer"
	"github.com/state-alchemists/zaruba/modules/strutil"
)

// Push monorepo and subtree
func Push(projectDir string) (err error) {
	if err = organizer.Organize(projectDir, organizer.NewOption()); err != nil {
		return err
	}
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return err
	}
	p, _, currentGitRemotes, err := git.LoadProjectConfig(projectDir)
	if err != nil {
		return err
	}
	// commit
	git.Commit(projectDir, fmt.Sprintf("Zaruba: Save before push to sub-repos at %s", time.Now().Format(time.RFC3339)))
	log.Println("[INFO] Push to main repo")
	if err = command.RunAndRedirect(projectDir, "git", "push", "origin", "HEAD"); err != nil {
		return err
	}
	subrepoPrefixMap := p.GetSubrepoPrefixMap(projectDir)
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		if !strutil.IsInArray(componentName, currentGitRemotes) {
			continue
		}
		component := p.GetComponentByName(componentName)
		location := component.GetLocation()
		origin := component.GetOrigin()
		branch := component.GetBranch()
		if location == "" || origin == "" || branch == "" {
			continue
		}
		log.Printf("[INFO] Push to sub-repo %s", componentName)
		if err = command.RunAndRedirect(projectDir, "git", "subtree", "push", "--prefix="+subrepoPrefix, componentName, branch); err != nil {
			return err
		}
	}
	return err
}
