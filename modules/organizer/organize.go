package organizer

import (
	"github.com/state-alchemists/zaruba/modules/action"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/file"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Organize projectDir
func Organize(projectDir string, p *config.ProjectConfig, arguments ...string) (err error) {
	sortedLinkSources := p.GetSortedLinkSources()
	links := p.GetLinks()
	ignores := p.GetIgnores()
	arguments = append([]string{projectDir}, arguments...)
	// pre-organize
	err = action.Do(
		"organize",
		action.CreateOption().
			SetWorkDir(projectDir).
			SetIsPerformAction(false).
			SetIsPerformPost(false).
			SetIgnores(ignores),
		arguments...,
	)
	if err != nil {
		return err
	}
	// copy
	for _, source := range sortedLinkSources {
		destinationList := links[source]
		err = copyLinks(source, destinationList)
		if err != nil {
			return err
		}
	}
	// organize and post-organize
	return action.Do(
		"organize",
		action.CreateOption().
			SetWorkDir(projectDir).
			SetIsPerformPre(false).
			SetIgnores(ignores),
		arguments...,
	)
}

func copyLinks(source string, destinationList []string) (err error) {
	// start multiple copyWithChannel as go-routines
	errChans := []chan error{}
	for _, destination := range destinationList {
		errChan := make(chan error)
		go copyWithChannel(source, destination, errChan)
		errChans = append(errChans, errChan)
	}
	// wait all go-routine finished
	for _, errChan := range errChans {
		err = <-errChan
		if err != nil {
			return err
		}
	}
	return err
}

func copyWithChannel(source, destination string, errChan chan error) {
	logger.Info("Copy `%s` to `%s`", source, destination)
	err := file.CopyExcept(source, destination, []string{`\.zaruba$`})
	errChan <- err
}
