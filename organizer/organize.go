package organizer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/state-alchemists/zaruba/action"
	"github.com/state-alchemists/zaruba/file"
	"github.com/state-alchemists/zaruba/format"
)

// Organize projectDir
func Organize(projectDir string, option *Option, arguments ...string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	log.Printf("[INFO] Organize project `%s` with option %s %s", projectDir, option.Sprintf(), format.SprintArgs(arguments))
	// remove depFile
	if err = os.Remove(filepath.Join(projectDir, "zaruba.dependency.json")); err != nil && !os.IsNotExist(err) {
		return
	}
	// link
	if err = action.Do("link", projectDir, action.NewOption()); err != nil {
		return
	}
	// get dep and sortedSources
	dep, sortedSources, err := getDepAndSort(projectDir)
	if err != nil {
		return
	}
	// update option.MTimeLimit
	for _, source := range sortedSources {
		var sourceMTime time.Time
		sourceMTime, err = file.GetMTime(source)
		if err != nil {
			return
		}
		destinationList := dep[source]
		for _, destination := range destinationList {
			if sourceMTime.Before(option.GetMTimeLimit()) {
				var destinationMTime time.Time
				destinationMTime, err = file.GetMTime(destination)
				if err != nil && os.IsNotExist(err) {
					option = getOptionBeforeSourceMTime(option, sourceMTime)
					log.Printf("[INFO] Update organizer.Option to %s because `%s` is not exists", option.Sprintf(), destination)
					break
				} else if destinationMTime.Before(sourceMTime) {
					option = getOptionBeforeSourceMTime(option, sourceMTime)
					log.Printf("[INFO] Update organizer.Option to %s because `%s` is older than `%s`", option.Sprintf(), destination, source)
					break
				}
			}
		}
	}
	return organize(projectDir, dep, sortedSources, option, arguments...)
}

func getOptionBeforeSourceMTime(option *Option, sourceMTime time.Time) *Option {
	return option.SetMTimeLimit(sourceMTime.Add(-time.Nanosecond))
}

func organize(projectDir string, dep map[string][]string, sortedSources []string, option *Option, arguments ...string) (err error) {
	// pre-organize
	err = action.Do(
		"organize-project", projectDir,
		action.NewOption().SetMTimeLimit(option.GetMTimeLimit()).SetPerformAction(false).SetPerformPre(false),
		arguments...,
	)
	// copy
	for _, source := range sortedSources {
		destinationList := dep[source]
		err = copyAll(option, source, destinationList)
		if err != nil {
			return
		}
	}
	// organize and post-organize
	err = action.Do(
		"organize-project", projectDir,
		action.NewOption().SetMTimeLimit(option.GetMTimeLimit()).SetPerformPre(false),
		arguments...,
	)
	return
}

func getDepAndSort(projectDir string) (dep map[string][]string, sortedSources []string, err error) {
	dep = map[string][]string{}
	sortedSources = []string{}
	// Read dependency file
	depFileName := filepath.Join(projectDir, "zaruba.dependency.json")
	jsonB, err := ioutil.ReadFile(depFileName)
	if err != nil {
		return
	}
	// unmarshal
	if err = json.Unmarshal(jsonB, &dep); err != nil {
		return
	}
	// get all keys of dep (i.e: list of sortedSources)
	for source := range dep {
		sortedSources = append(sortedSources, source)
	}
	// sort keys
	sort.SliceStable(sortedSources, func(i int, j int) bool {
		firstSource, secondSource := sortedSources[i], sortedSources[j]
		// get destination
		firstDestinations := dep[firstSource]
		// compare
		for _, destination := range firstDestinations {
			if strings.HasPrefix(destination, secondSource) {
				return true
			}
		}
		return false
	})
	return
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
			return
		}
	}
	return
}

func copyWithChannel(option *Option, source, destination string, errChan chan error) {
	sourceMTime, err := file.GetMTime(source)
	if err != nil {
		errChan <- err
		return
	}
	if sourceMTime.After(option.GetMTimeLimit()) {
		log.Printf("[INFO] Copy `%s` to `%s`", source, destination)
		err = file.CopyExcept(source, destination, []string{`\.zaruba$`})
	}
	errChan <- err
}
