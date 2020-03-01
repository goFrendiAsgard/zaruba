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
	if err = git.Init(projectDir); err != nil {
		return err
	}
	// get current branch name
	currentBranchName, err := git.GetCurrentBranchName(projectDir)
	if err != nil {
		return err
	}
	// get temporary branch name and checkout
	temporaryBranchName := getTemporaryBranchName()
	if err = git.Checkout(projectDir, temporaryBranchName, true); err != nil {
		return err
	}
	// get current git remotes
	currentGitRemotes, err := git.GetCurrentGitRemotes(projectDir)
	if err != nil {
		return err
	}
	// process subtree
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		if strutil.IsInArray(componentName, currentGitRemotes) {
			continue
		}
		if err = gitPullSubtree(p, projectDir, componentName, subrepoPrefix); err != nil {
			return err
		}
	}
	// checkout to current branch
	if err = git.Checkout(projectDir, currentBranchName, false); err != nil {
		return err
	}
	if err = git.Merge(projectDir, temporaryBranchName); err != nil {
		return err
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

func gitPullSubtree(p *config.ProjectConfig, projectDir, componentName, subrepoPrefix string) (err error) {
	component := p.Components[componentName]
	origin := component.Origin
	branch := component.Branch
	location := component.Location
	backupLocation := filepath.Join(projectDir, ".git", "subrepo-backup", subrepoPrefix)
	if location == "" || origin == "" || branch == "" {
		return nil
	}
	// backup
	log.Printf("[INFO] Prepare backup location")
	if err = os.MkdirAll(filepath.Dir(backupLocation), 0700); err != nil {
		return err
	}
	log.Printf("[INFO] Moving `%s` to `%s`", location, backupLocation)
	if err = os.Rename(location, backupLocation); err != nil {
		return err
	}
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
	if err := command.RunAndRedirect(projectDir, "git", "subtree", "add", "--prefix="+subrepoPrefix, componentName, branch); err != nil {
		return err
	}
	// fetch
	log.Printf("[INFO] Fetch `%s`", componentName)
	if err = command.RunAndRedirect(projectDir, "git", "fetch", componentName, branch); err != nil {
		return err
	}
	// restore
	log.Printf("[INFO] Restore")
	if err = os.RemoveAll(location); err != nil {
		return err
	}
	log.Printf("[INFO] Moving `%s` to `%s`", backupLocation, location)
	if err = os.Rename(backupLocation, location); err != nil {
		return err
	}
	// commit
	log.Printf("[INFO] Commit after subtree operation")
	git.Commit(projectDir, "after initiate "+componentName)
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
