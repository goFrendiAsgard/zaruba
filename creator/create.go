package creator

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
	"github.com/state-alchemists/zaruba/command"
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
	substitutions, err := getRealSubstitutions(modeConfig.Substitutions, isInteractive)
	if err != nil {
		return err
	}
	environ := os.Environ()
	log.Printf("Zaruba create target: %s", targetPath)
	if err = os.MkdirAll(targetPath, os.ModePerm); err != nil { // create target
		return err
	}
	log.Println("Zaruba run pre-triggers")
	if err = command.RunMultiple(shell, targetPath, environ, modeConfig.PreTriggers); err != nil { // run pre-triggers
		return err
	}
	log.Println("Zaruba perform copy")
	if err = copyToTarget(templatePath, targetPath, modeConfig.Copy); err != nil { // copy
		return err
	}
	log.Println("Zaruba perform copy and substitute")
	if err = copyToTargetAndSubstitute(templatePath, targetPath, substitutions, modeConfig.CopyAndSubstitute); err != nil { // copy
		return err
	}
	log.Println("Zaruba run post-triggers")
	if err = command.RunMultiple(shell, targetPath, environ, modeConfig.PostTriggers); err != nil { // run post-triggers
		return err
	}
	return err
}

func getRealSubstitutions(rawSubstitutions map[string]string, isInteractive bool) (substitutions map[string]string, err error) {
	substitutions = make(map[string]string)
	for key, val := range rawSubstitutions {
		if os.Getenv(key) != "" {
			val = os.Getenv(key)
		}
		if isInteractive {
			answer, err := getUserInput(fmt.Sprintf("Please set the value of `%s` (Default: `%s`)", key, val))
			if err != nil {
				return substitutions, err
			}
			if strings.Trim(answer, "\r\n") != "" {
				val = answer
			}
		}
		substitutions[key] = val
	}
	return
}

func getUserInput(text string) (answer string, err error) {
	reader := bufio.NewReader(os.Stdin)
	log.Println(text)
	answer, err = reader.ReadString('\n')
	return
}

func copyToTargetAndSubstitute(templatePath string, targetPath string, substitutions map[string]string, fileMap map[string]string) error {
	for src, dst := range fileMap {
		src := path.Join(templatePath, src)
		dst := path.Join(targetPath, dst)
		if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
			return err
		}
		// load template from src
		t := template.New(path.Base(src))
		t, err := t.ParseFiles(src)
		if err != nil {
			return err
		}
		// open dst file
		dstFile, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer dstFile.Close()
		// write to dst file
		err = t.Execute(dstFile, substitutions)
		if err != nil {
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
