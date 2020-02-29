package puller

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/organizer"
)

// Pull monorepo and subtree
func Pull(projectDir string) (err error) {
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
	command.RunAndRedirect(projectDir, "git", "commit", "-m", fmt.Sprintf("Save before pull on %s", time.Now().Format(time.RFC3339)))
	log.Println("[INFO] Pull from main repo")
	if err = command.RunAndRedirect(projectDir, "git", "pull", "origin", "HEAD"); err != nil {
		return err
	}
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
		log.Printf("[INFO] Pulling from sub-repo %s", componentName)
		command.RunAndRedirect(projectDir, "git", "subtree", "pull", "--prefix="+subrepoPrefix, "--squash", componentName, branch)
	}
	organizer.Organize(projectDir, organizer.NewOption())
	return err
}
