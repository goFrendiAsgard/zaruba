package pusher

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

// Push monorepo and subtree
func Push(projectDir string) (err error) {
	organizer.Organize(projectDir, organizer.NewOption())
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	log.Println("[INFO] Commit if there are changes")
	gitCommit(projectDir)
	log.Println("[INFO] Push repo")
	gitPush(projectDir)
	p := config.LoadProjectConfig(projectDir)
	for componentName, component := range p.Components {
		location := component.Location
		origin := component.Origin
		branch := component.Branch
		if location == "" || origin == "" || branch == "" {
			continue
		}
		log.Printf("[INFO] Pushing sub-repo %s", componentName)
		subRepoPrefix := getSubrepoPrefix(projectDir, location)
		var cmd *exec.Cmd
		cmd, err = command.GetShellCmd(projectDir, fmt.Sprintf(
			"git subtree push --prefix=%s %s %s",
			subRepoPrefix, componentName, branch,
		))
		if err != nil {
			return
		}
		command.Run(cmd)
	}
	return
}

func getSubrepoPrefix(projectDir, location string) string {
	if !strings.HasPrefix(location, projectDir) {
		return location
	}
	return strings.Trim(strings.TrimPrefix(location, projectDir), string(os.PathSeparator))
}

func gitCommit(projectDir string) (err error) {
	cmd, err := command.GetShellCmd(projectDir, fmt.Sprintf(
		"git add . -A && git commit -m 'Save before push on %s'",
		time.Now().Format(time.RFC3339),
	))
	if err != nil {
		return
	}
	return command.Run(cmd)
}

func gitPush(projectDir string) (err error) {
	cmd, err := command.GetShellCmd(projectDir, "git push origin HEAD")
	if err != nil {
		return
	}
	return command.Run(cmd)
}
