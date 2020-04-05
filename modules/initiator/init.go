package initiator

import (
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
	return organizer.Organize(projectDir, organizer.NewOption())
}

func gitProcessSubtree(p *config.ProjectConfig, projectDir, componentName, subrepoPrefix string) (err error) {
	component, err := p.GetComponentByName(componentName)
	if err != nil {
		return err
	}
	origin := component.GetOrigin()
	branch := component.GetBranch()
	location := component.GetLocation()
	backupLocation := filepath.Join(projectDir, ".git", ".subrepobackup", subrepoPrefix)
	if location == "" || origin == "" || branch == "" {
		return nil
	}
	// commit
	git.Commit(projectDir, fmt.Sprintf("Zaruba: prepare syncing `%s` on %s", componentName, time.Now().Format(time.RFC3339)))
	// backup
	if err = backup(location, backupLocation); err != nil {
		return err
	}
	// add remote
	logger.Info("Add remote `%s` as `%s`", origin, componentName)
	if err = command.RunAndRedirect(projectDir, "git", "remote", "add", componentName, origin); err != nil {
		return err
	}
	// commit
	git.Commit(projectDir, fmt.Sprintf("Zaruba: syncing `%s` at %s ", componentName, time.Now().Format(time.RFC3339)))
	// add subtree
	logger.Info("Add subtree `%s` with prefix `%s`", componentName, subrepoPrefix)
	if err := git.SubtreeAdd(projectDir, subrepoPrefix, componentName, branch); err != nil {
		return err
	}
	// fetch
	logger.Info("Fetch `%s`", componentName)
	if err = command.RunAndRedirect(projectDir, "git", "fetch", componentName, branch); err != nil {
		return err
	}
	if err = command.RunAndRedirect(projectDir, "git", "pull", componentName, branch); err != nil {
		return err
	}
	// restore
	if err := restore(backupLocation, location); err != nil {
		return err
	}
	// commit
	git.Commit(projectDir, fmt.Sprintf("Zaruba: successfully syncing `%s` on %s", componentName, time.Now().Format(time.RFC3339)))
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
