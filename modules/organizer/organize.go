package organizer

import (
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/file"
	"github.com/state-alchemists/zaruba/modules/logger"
)

// Organize projectDir
func Organize(projectDir string, p *config.ProjectConfig, arguments ...string) (err error) {
	sortedLinkSources := p.GetSortedLinkSources()
	links := p.GetLinks()
	arguments = append([]string{projectDir}, arguments...)
	// copy
	for _, source := range sortedLinkSources {
		destinationList := links[source]
		err = copyLinks(source, destinationList)
		if err != nil {
			return err
		}
	}
	return err
}

func copyLinks(source string, destinationList []string) (err error) {
	// start multiple copyWithChannel as go-routines
	errChan := make(chan error, len(destinationList))
	for _, destination := range destinationList {
		go copyWithChannel(source, destination, errChan)
	}
	// wait all go-routine finished
	for range destinationList {
		err = <-errChan
		if err != nil {
			return err
		}
	}
	return err
}

func copyWithChannel(source, destination string, errChan chan error) {
	logger.Info("📝 Copy `%s` to `%s`", source, destination)
	err := file.CopyExcept(source, destination, []string{`\.zaruba$`})
	errChan <- err
}
