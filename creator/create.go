package creator

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/otiai10/copy"
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
	environ := os.Environ()
	substitutions := getRealSubstitutions(modeConfig.Substitutions, isInteractive)
	log.Printf("Create target: %s", targetPath)
	if err = os.MkdirAll(targetPath, os.ModePerm); err != nil { // create target
		return err
	}
	log.Println("Run pre-triggers")
	if err = runMultipleCommands(shell, targetPath, environ, modeConfig.PreTriggers); err != nil { // run pre-triggers
		return err
	}
	log.Println("Copy")
	if err = copyToTarget(templatePath, targetPath, modeConfig.Copy); err != nil { // copy
		return err
	}
	log.Println("Copy and substitute")
	if err = copyToTargetAndSubstitute(templatePath, targetPath, substitutions, modeConfig.CopyAndSubstitute); err != nil { // copy
		return err
	}
	log.Println("Run post-triggers")
	if err = runMultipleCommands(shell, targetPath, environ, modeConfig.PostTriggers); err != nil { // run post-triggers
		return err
	}
	return err
}

func getRealSubstitutions(rawSubstitutions map[string]string, isInteractive bool) map[string]string {
	substitutions := make(map[string]string)
	for key, val := range rawSubstitutions {
		if os.Getenv(key) != "" {
			val = os.Getenv(key)
		}
		substitutions[key] = val
	}
	return substitutions
}

func copyToTargetAndSubstitute(templatePath string, targetPath string, substitutions map[string]string, fileMap map[string]string) error {
	for src, dst := range fileMap {
		src := path.Join(templatePath, src)
		dst := path.Join(targetPath, dst)
		if err := copy.Copy(src, dst); err != nil {
			return err
		}
	}
	return nil
}

func copyToTarget(templatePath string, targetPath string, fileMap map[string]string) error {
	for src, dst := range fileMap {
		src := path.Join(templatePath, src)
		dst := path.Join(targetPath, dst)
		if err := copy.Copy(src, dst); err != nil {
			return err
		}
	}
	return nil
}

func runMultipleCommands(shell []string, dir string, environ []string, commands []string) error {
	for _, command := range commands {
		err := runSingleCommand(shell, dir, environ, command)
		if err != nil {
			return err
		}
	}
	return nil
}

func runSingleCommand(shell []string, dir string, environ []string, command string) error {
	commandList := append(shell, command)
	cmd := exec.Command(commandList[0], commandList[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = environ
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
