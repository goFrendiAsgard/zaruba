package git

import (
	"log"

	"github.com/state-alchemists/zaruba/modules/config"
)

// LoadProjectConfig check project component origin
func LoadProjectConfig(projectDir string) (p *config.ProjectConfig, currentBranchName string, currentGitRemotes []string, err error) {
	// load project config
	log.Printf("[INFO] Load project config from `%s`", projectDir)
	p, err = config.LoadProjectConfig(projectDir)
	if err != nil {
		return p, currentBranchName, currentGitRemotes, err
	}
	// git init
	if err = Init(projectDir); err != nil {
		return p, currentBranchName, currentGitRemotes, err
	}
	// check all component's origin
	for componentName, component := range p.Components {
		origin := component.Origin
		if origin == "" {
			continue
		}
		log.Printf("[INFO] Check origin of component `%s`", componentName)
		if err = LsRemote(projectDir, origin); err != nil {
			return p, currentBranchName, currentGitRemotes, err
		}
	}
	// get currentBranchName
	currentBranchName, err = GetCurrentBranchName(projectDir)
	if err != nil {
		return p, currentBranchName, currentGitRemotes, err
	}
	// get currentGitRemotes
	currentGitRemotes, err = GetCurrentGitRemotes(projectDir)
	return p, currentBranchName, currentGitRemotes, err
}
