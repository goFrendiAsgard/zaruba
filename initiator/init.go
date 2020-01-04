package initiator

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/state-alchemists/zaruba/command"
	"github.com/state-alchemists/zaruba/config"
)

// Init monorepo and subtree
func Init(projectDir string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	log.Println("[INFO] Initiating master repo")
	gitInitMaster(projectDir)
	p := config.LoadProjectConfig(projectDir)
	for componentName, component := range p.Components {
		log.Printf("[INFO] Checking %s", componentName)
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
		var cmd *exec.Cmd
		cmd, err = command.GetShellCmd(projectDir, fmt.Sprintf(
			"git remote add %s %s && git subtree add --prefix=%s %s %s && git fetch %s %s",
			componentName, origin, // git remote add
			location, componentName, branch, // git subtree add
			componentName, branch, // git fetch
		))
		if err != nil {
			return
		}
		command.Run(cmd)
	}
	return
}

func gitInitMaster(projectDir string) (err error) {
	cmd, err := command.GetShellCmd(projectDir, "git init && git add . -A && git commit -m 'First Commmit'")
	if err != nil {
		return
	}
	return command.Run(cmd)
}
