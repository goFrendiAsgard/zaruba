package organizer

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/state-alchemists/zaruba/modules/action"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/file"
	"github.com/state-alchemists/zaruba/modules/strutil"
)

// Organize projectDir
func Organize(projectDir string, option *Option, arguments ...string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Organize project `%s` with option %s %s", projectDir, option.Sprintf(), strutil.SprintArgs(arguments))
	p, err := config.NewProjectConfig(projectDir)
	if err != nil {
		return err
	}
	str, _ := p.ToYaml()
	log.Printf("[INFO] Project Config Loaded:\n%s", str)
	sortedLinkSources := p.GetSortedLinkSources()
	// update option.MTimeLimit
	for _, source := range sortedLinkSources {
		var sourceMTime time.Time
		sourceMTime, err = file.GetMTime(source)
		if err != nil {
			return err
		}
		destinationList := p.GetLinkDestinationList(source)
		for _, destination := range destinationList {
			if sourceMTime.Before(option.GetMTimeLimit()) {
				var destinationMTime time.Time
				destinationMTime, err = file.GetMTime(destination)
				if err != nil && os.IsNotExist(err) {
					updateOptionToPreeceedSource(option, sourceMTime)
					log.Printf("[INFO] Update organizer.Option to %s because `%s` is not exists", option.Sprintf(), destination)
					break
				} else if destinationMTime.Before(sourceMTime) {
					updateOptionToPreeceedSource(option, sourceMTime)
					log.Printf("[INFO] Update organizer.Option to %s because `%s` is older than `%s`", option.Sprintf(), destination, source)
					break
				}
			}
		}
	}
	return organize(projectDir, p.GetLinks(), sortedLinkSources, option, arguments...)
}

func updateOptionToPreeceedSource(option *Option, sourceMTime time.Time) (updatedOption *Option) {
	return option.SetMTimeLimit(sourceMTime.Add(-time.Nanosecond))
}

func organize(projectDir string, links map[string][]string, sortedLinkSources []string, option *Option, arguments ...string) (err error) {
	arguments = append([]string{projectDir}, arguments...)
	// pre-organize
	err = action.Do(
		"organize",
		action.NewOption().
			SetWorkDir(projectDir).
			SetMTimeLimit(option.GetMTimeLimit()).
			SetIsPerformAction(false).
			SetIsPerformPost(false),
		arguments...,
	)
	if err != nil {
		return err
	}
	// copy
	for _, source := range sortedLinkSources {
		destinationList := links[source]
		err = copyAll(option, source, destinationList)
		if err != nil {
			return err
		}
	}
	// organize and post-organize
	return action.Do(
		"organize",
		action.NewOption().
			SetWorkDir(projectDir).
			SetMTimeLimit(option.GetMTimeLimit()).
			SetIsPerformPre(false),
		arguments...,
	)
}

func copyAll(option *Option, source string, destinationList []string) (err error) {
	// start multiple copyWithChannel as go-routines
	errChans := []chan error{}
	for _, destination := range destinationList {
		errChan := make(chan error)
		go copyWithChannel(option, source, destination, errChan)
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

func copyWithChannel(option *Option, source, destination string, errChan chan error) {
	sourceMTime, err := file.GetMTime(source)
	if err != nil {
		errChan <- err
		return
	}
	if sourceMTime.After(option.GetMTimeLimit()) {
		log.Printf("[INFO] Copy `%s` to `%s`", source, destination)
		err = file.CopyExcept(source, destination, []string{
			`\.zaruba$`,
		})
	}
	errChan <- err
}
