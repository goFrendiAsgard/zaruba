package initiator

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
	"github.com/state-alchemists/zaruba/organizer"
)

// Init monorepo and subtree
func Init(projectDir string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	log.Println("[INFO] Initiating repo")
	gitInitAndCommit(projectDir)
	p := config.LoadProjectConfig(projectDir)
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
		var cmd *exec.Cmd
		cmd, err = command.GetShellCmd(projectDir, fmt.Sprintf(
			"git remote add %s %s && git subtree add --prefix=%s %s %s && git fetch %s %s",
			componentName, origin, // git remote add
			subRepoPrefix, componentName, branch, // git subtree add
			componentName, branch, // git fetch
		))
		if err != nil {
			return
		}
		command.Run(cmd)
	}
	organizer.Organize(projectDir, organizer.NewOption())
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
