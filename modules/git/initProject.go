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

// InitProject check project component origin
func InitProject(projectDir string, p *config.ProjectConfig) (err error) {
	// git init
	if err = Init(projectDir); err != nil {
		return err
	}
	// check all component's origin
	errChans := []chan error{}
	for componentName, component := range p.GetComponents() {
		errChan := make(chan error)
		errChans = append(errChans, errChan)
		go checkSubrepo(projectDir, componentName, component, errChan)
	}
	for _, errChan := range errChans {
		if err = <-errChan; err != nil {
			return err
		}
	}
	return err
}

func checkSubrepo(projectDir string, componentName string, component *config.Component, errChan chan error) {
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
		if err = initSubrepo(projectDir, componentName, component); err != nil {
			errChan <- err
			return
		}
	}
	errChan <- nil
	return
}

func initSubrepo(projectDir, componentName string, component *config.Component) (err error) {
	tempDir := filepath.Join(projectDir, ".git", ".newsubrepo", componentName)
	if err = os.MkdirAll(tempDir, 0700); err != nil {
		return err
	}
	origin := component.GetOrigin()
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
	Commit(tempDir, fmt.Sprintf("💀 First commit for `%s` at %s", componentName, time.Now().Format(time.RFC3339)))
	// add remote
	if err = command.RunAndRedirect(tempDir, "git", "remote", "add", "origin", origin); err != nil {
		return err
	}
	// push
	if err = command.RunAndRedirect(tempDir, "git", "push", "-u", "origin", "master"); err != nil {
		return err
	}
	return os.RemoveAll(tempDir)
}
