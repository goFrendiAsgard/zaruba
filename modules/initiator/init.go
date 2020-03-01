package initiator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/git"
	"github.com/state-alchemists/zaruba/modules/organizer"
	"github.com/state-alchemists/zaruba/modules/strutil"
)

// Init monorepo and subtree
func Init(projectDir string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return err
	}
	if err = createZarubaConfigIfNotExists(projectDir); err != nil {
		return err
	}
	// load project
	p, _, currentGitRemotes, err := git.LoadProjectConfig(projectDir)
	if err != nil {
		return err
	}
	// process subtree
	subrepoPrefixMap := p.GetSubrepoPrefixMap(projectDir)
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		if strutil.IsInArray(componentName, currentGitRemotes) {
			continue
		}
		if err = gitProcessSubtree(p, projectDir, componentName, subrepoPrefix); err != nil {
			command.RunAndRedirect(projectDir, "git", "remote", "remove", componentName)
			return err
		}
	}
	// organize
	return organizer.Organize(projectDir, organizer.NewOption())
}

func getTemporaryBranchName() (branchName string) {
	branchName = fmt.Sprintf("init-%s", time.Now().Format(time.RFC3339))
	branchName = strings.ReplaceAll(branchName, "+", "Z")
	branchName = strings.ReplaceAll(branchName, ":", "-")
	return branchName
}

func gitProcessSubtree(p *config.ProjectConfig, projectDir, componentName, subrepoPrefix string) (err error) {
	component := p.Components[componentName]
	origin := component.Origin
	branch := component.Branch
	location := component.Location
	backupLocation := filepath.Join(projectDir, ".git", "subrepo-backup", subrepoPrefix)
	if location == "" || origin == "" || branch == "" {
		return nil
	}
	// backup
	backup(location, backupLocation)
	// add remote
	log.Printf("[INFO] Add remote `%s` as `%s`", origin, componentName)
	if err = command.RunAndRedirect(projectDir, "git", "remote", "add", componentName, origin); err != nil {
		return err
	}
	// commit
	log.Printf("[INFO] Commit before subtree operation")
	git.Commit(projectDir, "prepare to initiate "+componentName)
	// add subtree
	log.Printf("[INFO] Add subtree `%s` with prefix `%s`", componentName, subrepoPrefix)
	if err := git.SubtreeAdd(projectDir, subrepoPrefix, componentName, branch); err != nil {
		return err
	}
	// fetch
	log.Printf("[INFO] Fetch `%s`", componentName)
	if err = command.RunAndRedirect(projectDir, "git", "fetch", componentName, branch); err != nil {
		return err
	}
	// restore
	restore(backupLocation, location)
	// commit
	log.Printf("[INFO] Commit after subtree operation")
	git.Commit(projectDir, "after initiate "+componentName)
	return err
}

func restore(backupLocation, location string) (err error) {
	// restore
	log.Printf("[INFO] Restore")
	if err = os.RemoveAll(location); err != nil {
		return err
	}
	log.Printf("[INFO] Moving `%s` to `%s`", backupLocation, location)
	return os.Rename(backupLocation, location)
}

func backup(location, backupLocation string) (err error) {
	// backup
	log.Printf("[INFO] Prepare backup location")
	os.RemoveAll(backupLocation)
	if err = os.MkdirAll(filepath.Dir(backupLocation), 0777); err != nil {
		return err
	}
	log.Printf("[INFO] Moving `%s` to `%s`", location, backupLocation)
	return os.Rename(location, backupLocation)
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
