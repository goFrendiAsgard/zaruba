package initiator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/state-alchemists/zaruba/modules/command"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/organizer"
)

// Init monorepo and subtree
func Init(projectDir string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	log.Println("[INFO] Initiating repo")
	gitInitAndCommit(projectDir)
	// create config file if not exists
	if err = createZarubaConfigIfNotExists(projectDir); err != nil {
		return
	}
	p, err := config.LoadProjectConfig(projectDir)
	if err != nil {
		return
	}
	for componentName, component := range p.Components {
		location := component.Location
		origin := component.Origin
		branch := component.Branch
		if location == "" || origin == "" || branch == "" {
			continue
		}
		locationDir := filepath.Dir(location)
		if _, statErr := os.Stat(locationDir); os.IsNotExist(statErr) {
			log.Printf("[INFO] Creating directory %s", locationDir)
			os.Mkdir(locationDir, os.ModePerm)
		}
		log.Printf("[INFO] Initiating sub-repo %s", componentName)
		subRepoPrefix := getSubrepoPrefix(projectDir, location)
		tmpLocation := filepath.Join(".git", "tmp", subRepoPrefix)
		var cmd *exec.Cmd
		cmd, err = command.GetShellCmd(projectDir, fmt.Sprintf(
			"(mkdir -p %s && mv %s %s && git remote add %s %s && git subtree add --prefix=%s %s %s && git fetch %s %s && mv %s %s) || mv %s %s",
			tmpLocation,                // mkdir -p
			subRepoPrefix, tmpLocation, // mv
			componentName, origin, // git remote add
			subRepoPrefix, componentName, branch, // git subtree add
			componentName, branch, // git fetch
			tmpLocation, subRepoPrefix, // mv
			tmpLocation, subRepoPrefix, // mv
		))
		if err != nil {
			return
		}
		command.Run(cmd)
	}
	organizer.Organize(projectDir, organizer.NewOption())
	return
}

func createZarubaConfigIfNotExists(projectDir string) (err error) {
	zarubaConfigFile := filepath.Join(projectDir, "zaruba.config.yaml")
	if _, statErr := os.Stat(zarubaConfigFile); os.IsNotExist(statErr) {
		var configYaml string
		configYaml, err = config.NewProjectConfig().ToYaml()
		if err != nil {
			return
		}
		log.Printf("[INFO] Write to %s:\n%s", zarubaConfigFile, configYaml)
		ioutil.WriteFile(zarubaConfigFile, []byte(configYaml), 0755)
	}
	return
}

func getSubrepoPrefix(projectDir, location string) string {
	if !strings.HasPrefix(location, projectDir) {
		return location
	}
	return strings.Trim(strings.TrimPrefix(location, projectDir), string(os.PathSeparator))
}

func gitInitAndCommit(projectDir string) (err error) {
	cmd, err := command.GetShellCmd(projectDir, fmt.Sprintf(
		"git init && git add . -A && git commit -m 'First commit on %s'",
		time.Now().Format(time.RFC3339),
	))
	if err != nil {
		return
	}
	return command.Run(cmd)
}
