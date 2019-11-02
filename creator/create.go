package creator

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/state-alchemists/zaruba/config"
)

// Create something from template
func Create(template string, targetPath string, isInteractive bool) error {
	templateDir := config.GetTemplateDir()
	shell := config.GetShell()
	templatePath, mode := getTemplatePathAndMode(templateDir, template)
	modeConfig, err := getModeConfig(templatePath, mode)
	if err != nil {
		return err
	}
	log.Printf("%#v", modeConfig)
	log.Println(shell)
	log.Println("Creating target")
	if err = os.MkdirAll(targetPath, os.ModePerm); err != nil { // create target
		return err
	}
	log.Println("Running pre-triggers")
	if err = runMultipleCommands(shell, targetPath, modeConfig.PreTriggers); err != nil { // run pre-triggers
		return err
	}
	log.Println("Running post-triggers")
	if err = runMultipleCommands(shell, targetPath, modeConfig.PostTriggers); err != nil { // run post-triggers
		return err
	}
	return err
}

func runMultipleCommands(shell []string, dir string, commands []string) error {
	for _, command := range commands {
		err := runSingleCommand(shell, dir, command)
		if err != nil {
			return err
		}
	}
	return nil
}

func runSingleCommand(shell []string, dir string, command string) error {
	commandList := append(shell, command)
	cmd := exec.Command(commandList[0], commandList[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return err
	}
	return err
}

func getModeConfig(templatePath string, mode string) (modeConfig ModeConfig, err error) {
	templateConfig, err := NewTemplateConfig(templatePath)
	if err != nil {
		return
	}
	modeConfig, err = templateConfig.GetMode(mode)
	return
}

func getTemplatePathAndMode(templateDir string, template string) (templatePath string, mode string) {
	rawTemplatePath, mode := getRawTemplatePathAndMode(template)
	templatePath = path.Join(templateDir, rawTemplatePath)
	return
}

func getRawTemplatePathAndMode(template string) (rawTemplatePath string, mode string) {
	parts := strings.Split(template, ":")
	rawTemplatePath = parts[0]
	if len(parts) == 1 {
		mode = "base"
		return
	}
	mode = strings.Join(parts[1:], ":")
	return
}
