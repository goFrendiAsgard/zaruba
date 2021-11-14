package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install external tools",
}

func AddInstallCmdSubCommand() {
	setupFilePath := filepath.Join(os.Getenv("ZARUBA_HOME"), "setup")
	util := core.NewUtil()
	fileList, err := util.File.ListDir(setupFilePath)
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
				decoration := output.NewDecoration()
				logger := output.NewConsoleLogger(decoration)
				if err := runInstallSubCmd(shell, filepath.Join(setupFilePath, fileName)); err != nil {
					exit(cmd, logger, decoration, err)
				}
			},
		}
		installCmd.AddCommand(subCommand)
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
