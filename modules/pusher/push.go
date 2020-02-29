package pusher

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/organizer"
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
	log.Println("[INFO] Commit if there are changes")
	if err = command.RunAndRedirect(projectDir, "git", "add", ".", "-A"); err != nil {
		return err
	}
	command.RunAndRedirect(projectDir, "git", "commit", "-m", fmt.Sprintf("Save before push on %s", time.Now().Format(time.RFC3339)))
	log.Println("[INFO] Push to main repo")
	if err = command.RunAndRedirect(projectDir, "git", "push", "origin", "HEAD"); err != nil {
		return err
	}
	// get project config
	p, err := config.LoadProjectConfig(projectDir)
	if err != nil {
		return err
	}
	subrepoPrefixMap := p.GetSubrepoPrefixMap(projectDir)
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		component := p.Components[componentName]
		location := component.Location
		origin := component.Origin
		branch := component.Branch
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
