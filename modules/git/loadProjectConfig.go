package git

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// LoadProjectConfig check project component origin
func LoadProjectConfig(projectDir string) (p *config.ProjectConfig, currentBranchName string, currentGitRemotes []string, err error) {
	// load project config
	logger.Info("Load project config from `%s`", projectDir)
	p, err = config.NewProjectConfig(projectDir)
	if err != nil {
		return p, currentBranchName, currentGitRemotes, err
	}
	// git init
	if err = Init(projectDir); err != nil {
		return p, currentBranchName, currentGitRemotes, err
	}
	// check all component's origin
	errChans := []chan error{}
	for componentName, component := range p.GetComponents() {
		errChan := make(chan error)
		errChans = append(errChans, errChan)
		go checkComponent(projectDir, componentName, component, errChan)
	}
	for _, errChan := range errChans {
		if err = <-errChan; err != nil {
			return p, currentBranchName, currentGitRemotes, err
		}
	}
	// get currentBranchName
	currentBranchName, getCurrentBranchErr := GetCurrentBranchName(projectDir)
	if getCurrentBranchErr != nil {
		currentBranchName = "master"
	}
	// get currentGitRemotes
	currentGitRemotes, err = GetCurrentGitRemotes(projectDir)
	return p, currentBranchName, currentGitRemotes, err
}

func checkComponent(projectDir string, componentName string, component *config.Component, errChan chan error) {
	origin := component.GetOrigin()
	if origin == "" {
		errChan <- nil
		return
	}
	logger.Info("Check origin of component `%s`", componentName)
	output, err := LsRemote(projectDir, origin)
	if err != nil {
		errChan <- err
		return
	}
	// an empty repo doesn't output anything.
	if output == "" {
		if err = initRepo(projectDir, componentName, component); err != nil {
			errChan <- err
			return
		}
	}
	errChan <- nil
	return
}

func initRepo(projectDir, componentName string, component *config.Component) (err error) {
	tempDir := filepath.Join(projectDir, ".git", ".newsubrepo", componentName)
	if err = os.MkdirAll(tempDir, 0700); err != nil {
		return err
	}
	origin := component.GetOrigin()
	branch := component.GetBranch()
	// init
	if err = Init(tempDir); err != nil {
		return err
	}
	// create README.md
	f, err := os.Create(filepath.Join(tempDir, "README.md"))
	if err != nil {
		return err
	}
	f.WriteString("# " + componentName)
	// commit
	Commit(tempDir, fmt.Sprintf("Zaruba: First commit for `%s` at %s", componentName, time.Now().Format(time.RFC3339)))
	// add remote
	if err = command.RunAndRedirect(tempDir, "git", "remote", "add", "origin", origin); err != nil {
		return err
	}
	// checkout if branch != master
	if branch != "master" {
		if err = Checkout(tempDir, branch, true); err != nil {
			return err
		}
	}
	// push
	if err = command.RunAndRedirect(tempDir, "git", "push", "-u", "origin", branch); err != nil {
		return err
	}
	return os.RemoveAll(tempDir)
}
