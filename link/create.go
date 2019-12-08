package link

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

// Create dependency link
func Create(projectDir, source, destination string) (err error) {
	// define paths and dep
	if projectDir, err = filepath.Abs(projectDir); err != nil {
		return
	}
	if source, err = filepath.Abs(source); err != nil {
		return
	}
	if destination, err = filepath.Abs(destination); err != nil {
		return
	}
	depFileName := filepath.Join(projectDir, "zaruba.dependency.json")
	// create `depFileName` if it is not exists
	if _, err := os.Stat(depFileName); os.IsNotExist(err) {
		os.Create(depFileName)
		ioutil.WriteFile(depFileName, []byte("{}"), 0644)
	}
	// open `depFile`
	depFile, err := os.Open(depFileName)
	if err != nil {
		defer depFile.Close()
	}
	// lock `depFile`
	if err = syscall.Flock(int(depFile.Fd()), syscall.LOCK_EX); err != nil {
		return
	}
	// add dependency
	if err = addDependency(depFileName, source, destination); err != nil {
		return
	}
	// unloack `depFile`
	err = syscall.Flock(int(depFile.Fd()), syscall.LOCK_UN)
	return
}

func addDependency(depFileName, source, destination string) (err error) {
	// read `dep` from `defFileName`
	b, err := ioutil.ReadFile(depFileName)
	if err != nil {
		return
	}
	dep := map[string][]string{}
	if err = json.Unmarshal(b, &dep); err != nil {
		dep = map[string][]string{} // encounter invalid JSON, assume the file was empty and continue
		log.Printf("[WARNING] Invalid JSON format: %s", err)
	}
	// add `source` and `destination` to `dep`
	if _, sourceExists := dep[source]; !sourceExists {
		dep[source] = []string{}
	}
	dep[source] = append(dep[source], destination)
	// write `dep` to `depFileName`, assuming operation should always success
	if b, err = json.Marshal(dep); err != nil {
		return
	}
	err = ioutil.WriteFile(depFileName, b, 0644)
	return
}
