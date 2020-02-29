package initiator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

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
	temporaryBranchName := fmt.Sprintf("sync-%s", time.Now().Format(time.RFC3339))
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
	command.RunScript(projectDir, "git init")
	// git checkout
	log.Printf("[INFO] Checkout to `%s`", temporaryBranchName)
	command.RunScript(projectDir, fmt.Sprintf("git checkout -b %s", temporaryBranchName))
	// backup all services
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		location := p.Components[componentName].Location
		locationDir := filepath.Dir(location)
		// create locationDir in case of it doesn't exists. We need it for git subtree add
		if _, statErr := os.Stat(locationDir); os.IsNotExist(statErr) {
			log.Printf("[INFO] Creating directory `%s`", locationDir)
			os.Mkdir(locationDir, os.ModePerm)
		}
		if _, statErr := os.Stat(location); os.IsNotExist(statErr) {
			continue
		}
		backupLocation := filepath.Join(projectDir, ".git", ".backup", subrepoPrefix)
		os.MkdirAll(filepath.Dir(backupLocation), 0700)
		log.Printf("[INFO] Moving `%s` to `%s`", location, backupLocation)
		os.Rename(location, backupLocation)
	}
	// git commit before add subtrees
	log.Println("[INFO] Commit changes before adding subtrees")
	command.RunScript(projectDir, "git add . -A && git commit -m 'Remove services from monorepo'")
	// add subtree
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		component := p.Components[componentName]
		origin := component.Origin
		branch := component.Branch
		command.RunScript(projectDir, fmt.Sprintf("git remote remove %s", componentName))
		command.RunScript(projectDir, fmt.Sprintf("git remote add %s %s", componentName, origin))
		command.RunScript(projectDir, fmt.Sprintf("git subtree add --prefix=%s %s %s", subrepoPrefix, componentName, branch))
		command.RunScript(projectDir, fmt.Sprintf("git fetch %s %s", componentName, branch))
	}
	// restore all services
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		location := p.Components[componentName].Location
		backupLocation := filepath.Join(projectDir, ".git", ".backup", subrepoPrefix)
		if _, statErr := os.Stat(backupLocation); os.IsNotExist(statErr) {
			continue
		}
		log.Printf("[INFO] Moving `%s` to `%s`", backupLocation, location)
		os.RemoveAll(location)
		os.Rename(backupLocation, location)
	}
	// git commit before add subtrees
	log.Println("[INFO] Commit changes before adding subtrees")
	command.RunScript(projectDir, "git add . -A && git commit -m 'Re-add services to monorepo'")
	// git checkout master
	log.Printf("[INFO] Checkout to master and merge `%s`", temporaryBranchName)
	command.RunScript(projectDir, fmt.Sprintf("git checkout master && git merge --squash  %s", temporaryBranchName))
	// organize
	organizer.Organize(projectDir, organizer.NewOption())
	return err
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
