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
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return err
	}
	log.Println("[INFO] Commit if there are changes")
	command.RunScript(projectDir, fmt.Sprintf(
		"git add . -A && git commit -m 'Save before pull on %s'",
		time.Now().Format(time.RFC3339),
	))
	log.Println("[INFO] Pull repo")
	command.RunScript(projectDir, "git pull origin HEAD")
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
		log.Printf("[INFO] Pulling sub-repo %s", componentName)
		command.RunScript(projectDir, fmt.Sprintf(
			"git subtree pull --prefix=%s --squash %s %s",
			subrepoPrefix, componentName, branch,
		))
	}
	organizer.Organize(projectDir, organizer.NewOption())
	return err
}
