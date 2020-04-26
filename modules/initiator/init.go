package initiator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/git"
	"github.com/state-alchemists/zaruba/modules/logger"
	"github.com/state-alchemists/zaruba/modules/organizer"
	"github.com/state-alchemists/zaruba/modules/strutil"
)

// Init monorepo and subtree
func Init(projectDir string, p *config.ProjectConfig) (err error) {
	if branchName, err := git.GetCurrentBranchName(projectDir); err == nil {
		if branchName != "master" {
			return errors.New("You are not on `master` branch, please checkout to `master` and continue")
		}
	}
	if err = createZarubaConfigIfNotExists(projectDir); err != nil {
		return err
	}
	// Init project
	if err = git.InitProject(projectDir, p); err != nil {
		return err
	}
	// get remotes
	currentGitRemotes, err := git.GetCurrentGitRemotes(projectDir)
	if err != nil {
		return err
	}
	// get valid subrepo prefix map
	validSubrepoPrefixMap, err := getValidSubrepoPrefixMap(p, projectDir, currentGitRemotes)
	if err != nil {
		return err
	}
	for componentName, subrepoPrefix := range validSubrepoPrefixMap {
		if err = gitProcessSubtree(p, projectDir, componentName, subrepoPrefix); err != nil {
			command.RunAndRedirect(projectDir, "git", "remote", "remove", componentName)
			logger.Error("%s", err)
			break
		}
	}
	// organize
	return organizer.Organize(projectDir, p)
}

func gitProcessSubtree(p *config.ProjectConfig, projectDir, componentName, subrepoPrefix string) (err error) {
	component, err := p.GetComponentByName(componentName)
	if err != nil {
		return err
	}
	branch, err := git.GetCurrentBranchName(projectDir)
	if err != nil {
		return err
	}
	origin := component.GetOrigin()
	location := component.GetLocation()
	backupLocation := filepath.Join(projectDir, ".git", ".subrepobackup", subrepoPrefix)
	if location == "" || origin == "" {
		return nil
	}
	// commit
	git.CommitIfAnyDiff(projectDir, fmt.Sprintf("💀 Backup local %s at: %s", componentName, time.Now().Format(time.RFC3339)))
	// backup
	if err = backup(location, backupLocation); err != nil {
		restore(backupLocation, location)
		return err
	}
	// add remote
	if err = command.RunAndRedirect(projectDir, "git", "remote", "add", componentName, origin); err != nil {
		restore(backupLocation, location)
		return err
	}
	// commit
	git.CommitIfAnyDiff(projectDir, fmt.Sprintf("💀 Remove local %s at: %s", componentName, time.Now().Format(time.RFC3339)))
	// add subtree
	if err := git.SubtreeAdd(projectDir, subrepoPrefix, componentName, branch); err != nil {
		restore(backupLocation, location)
		return err
	}
	// fetch
	if err = command.RunAndRedirect(projectDir, "git", "fetch", componentName, branch); err != nil {
		restore(backupLocation, location)
		return err
	}
	if err = command.RunAndRedirect(projectDir, "git", "pull", componentName, branch); err != nil {
		restore(backupLocation, location)
		return err
	}
	// restore
	if err := restore(backupLocation, location); err != nil {
		return err
	}
	// commit
	git.CommitIfAnyDiff(projectDir, fmt.Sprintf("💀 Restore local %s at: %s", componentName, time.Now().Format(time.RFC3339)))
	return err
}

func restore(backupLocation, location string) (err error) {
	// restore
	logger.Info("Restore")
	if err = os.RemoveAll(location); err != nil {
		return err
	}
	logger.Info("Moving `%s` to `%s`", backupLocation, location)
	return os.Rename(backupLocation, location)
}

func backup(location, backupLocation string) (err error) {
	// backup
	logger.Info("Prepare backup location")
	os.RemoveAll(backupLocation)
	if err = os.MkdirAll(filepath.Dir(backupLocation), 0700); err != nil {
		return err
	}
	logger.Info("Moving `%s` to `%s`", location, backupLocation)
	return os.Rename(location, backupLocation)
}

func createZarubaConfigIfNotExists(projectDir string) (err error) {
	zarubaConfigFile := filepath.Join(projectDir, "zaruba.config.yaml")
	if _, statErr := os.Stat(zarubaConfigFile); os.IsNotExist(statErr) {
		p, err := config.NewProjectConfig(projectDir)
		if err != nil {
			return err
		}
		configYaml, err := p.ToYaml()
		if err != nil {
			return err
		}
		logger.Info("Write to %s:\n%s", zarubaConfigFile, configYaml)
		ioutil.WriteFile(zarubaConfigFile, []byte(configYaml), 0755)
	}
	return
}

func getValidSubrepoPrefixMap(p *config.ProjectConfig, projectDir string, currentGitRemotes []string) (validSubRepoPrefixMap map[string]string, err error) {
	// process subtree
	subrepoPrefixMap := p.GetSubrepoPrefixMap(projectDir)
	validSubRepoPrefixMap = map[string]string{}
	for componentName, subrepoPrefix := range subrepoPrefixMap {
		if strutil.IsInArray(componentName, currentGitRemotes) {
			continue
		}
		validSubRepoPrefixMap[componentName] = subrepoPrefix
	}
	return validSubRepoPrefixMap, err
}
