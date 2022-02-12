package installcmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var Cmd = &cobra.Command{
	Use:   "install",
	Short: "Install external tools",
}

func Init() {
	setupFilePath := filepath.Join(os.Getenv("ZARUBA_HOME"), "setup")
	util := core.NewCoreUtil()
	fileList, err := util.File.List(setupFilePath)
	if err != nil {
		fmt.Printf("warning: %s is not found", setupFilePath)
	}
	shell := os.Getenv("ZARUBA_SHELL")
	for _, f := range fileList {
		fileName := f
		if !strings.HasSuffix(fileName, ".sh") {
			continue
		}
		fileNameParts := strings.Split(fileName, ".")
		subCommandName := fileNameParts[len(fileNameParts)-2]
		subCommand := &cobra.Command{
			Use: subCommandName,
			Run: func(cmd *cobra.Command, args []string) {
				decoration := output.NewDefaultDecoration()
				logger := output.NewConsoleLogger(decoration)
				if err := runInstallSubCmd(shell, filepath.Join(setupFilePath, fileName)); err != nil {
					cmdHelper.Exit(cmd, args, logger, decoration, err)
				}
			},
		}
		Cmd.AddCommand(subCommand)
	}

}

func runInstallSubCmd(shell, script string) error {
	fmt.Println(script)
	cmd := exec.Command(shell, "-c", script)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
