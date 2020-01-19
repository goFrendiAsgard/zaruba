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
	organizer.Organize(projectDir, organizer.NewOption())
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	log.Println("[INFO] Commit if there are changes")
	command.RunScript(projectDir, fmt.Sprintf(
		"git add . -A && git commit -m 'Save before push on %s'",
		time.Now().Format(time.RFC3339),
	))
	log.Println("[INFO] Push repo")
	command.RunScript(projectDir, "git push origin HEAD")
	p, err := config.LoadProjectConfig(projectDir)
	if err != nil {
		return
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
		log.Printf("[INFO] Pushing sub-repo %s", componentName)
		command.RunScript(projectDir, fmt.Sprintf(
			"git subtree push --prefix=%s %s %s",
			subrepoPrefix, componentName, branch,
		))
	}
	return
}
