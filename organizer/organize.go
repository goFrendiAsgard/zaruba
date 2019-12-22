package organizer

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/otiai10/copy"
	"github.com/state-alchemists/zaruba/action"
	"github.com/state-alchemists/zaruba/file"
)

// Organize projectDir
func Organize(projectDir string, option *Option, arguments ...string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
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
		err = copy.Copy(source, destination)
	}
	errChan <- err
}
