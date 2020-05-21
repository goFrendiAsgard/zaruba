package cmd

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/initiator"
	"github.com/state-alchemists/zaruba/modules/logger"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [project-dir]",
	Short: "Init a project",
	Long:  `Zaruba will init a project`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir := "."
		if len(args) > 0 {
			projectDir = args[0]
		}
		// get absolute project dir and project config
		projectDir, err := filepath.Abs(projectDir)
		if err != nil {
			logger.Fatal(err)
		}
		p, err := config.CreateProjectConfig(projectDir)
		if err != nil {
			logger.Fatal(err)
		}
		// listen to kill signal
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		// invoke action
		if err := initiator.Init(projectDir, p); err != nil {
			logger.Fatal(err)
		}
		// Don't handle kill signal unless the action was finished
		go func() {
			<-c
		}()
	},
}
