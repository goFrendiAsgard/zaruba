package initiator

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/organizer"
)

// Init monorepo and subtree
func Init(projectDir string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return err
	}
	// load project config
	log.Println("[INFO] Load Project Config")
	if err = createZarubaConfigIfNotExists(projectDir); err != nil {
		return err
	}
	p, err := config.LoadProjectConfig(projectDir)
	if err != nil {
		return err
	}
	subrepoPrefixMap := p.GetSubrepoPrefixMap(projectDir)
	// git init
	log.Println("[INFO] Initiating repo")
	if err = command.RunAndRedirect(projectDir, "git", "init"); err != nil {
		return err
	}
	// get current git remotes
	currentGitRemotes, err := getCurrentGitRemotes(projectDir)
	if err != nil {
		return err
	}
	// add subtree
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		if isComponentNameInCurrentGitRemotes(componentName, currentGitRemotes) {
			continue
		}
		component := p.Components[componentName]
		origin := component.Origin
		branch := component.Branch
		if err = command.RunAndRedirect(projectDir, "git", "remote", "add", componentName, origin); err != nil {
			return err
		}
		if err = command.RunAndRedirect(projectDir, "git", "subtree", "add", "--prefix="+subrepoPrefix, componentName, branch); err != nil {
			return err
		}
		if err = command.RunAndRedirect(projectDir, "git", "fetch", componentName, branch); err != nil {
			return err
		}
	}
	// organize
	return organizer.Organize(projectDir, organizer.NewOption())
}

func isComponentNameInCurrentGitRemotes(componentName string, currentGitRemotes []string) (exists bool) {
	for _, remote := range currentGitRemotes {
		if remote == componentName {
			return true
		}
	}
	return false
}

func getCurrentGitRemotes(projectDir string) (currentGitRemotes []string, err error) {
	log.Println("[INFO] Get current remotes")
	output, err := command.Run(projectDir, "git", "remote")
	if err != nil {
		return currentGitRemotes, err
	}
	outputList := strings.Split(output, "\n")
	for _, remote := range outputList {
		remote = strings.Trim(remote, "\r\n ")
		if remote != "" {
			log.Printf("[INFO] * %s", remote)
			currentGitRemotes = append(currentGitRemotes, remote)
		}
	}
	return currentGitRemotes, err
}

func createZarubaConfigIfNotExists(projectDir string) (err error) {
	zarubaConfigFile := filepath.Join(projectDir, "zaruba.config.yaml")
	if _, statErr := os.Stat(zarubaConfigFile); os.IsNotExist(statErr) {
		var configYaml string
		configYaml, err = config.NewProjectConfig().ToYaml()
		if err != nil {
			return
		}
		log.Printf("[INFO] Write to %s:\n%s", zarubaConfigFile, configYaml)
		ioutil.WriteFile(zarubaConfigFile, []byte(configYaml), 0755)
	}
	return
}
