package cmd

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/modules/config"
	"github.com/state-alchemists/zaruba/modules/logger"
	"github.com/state-alchemists/zaruba/modules/runner"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run [project-dir [selector-1 ... selector-N]]",
	Short: "Run project.",
	Long:  "Run components defined in a project",
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir, err := filepath.Abs(".")
		if err != nil {
			logger.Fatal(err)
		}
		p, err := config.CreateProjectConfig(projectDir)
		if err != nil {
			logger.Fatal(err)
		}
		// invoke action
		stopChan := make(chan bool)
		executedChan := make(chan bool)
		errChan := make(chan error)
		// listen to kill signal
		osSignalChan := make(chan os.Signal)
		signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
		go func() {
			<-osSignalChan
			stopChan <- true
			time.Sleep(60 * time.Second) // after one minute if the process is not stopped, kill it
			logger.Fatal("Cannot terminate process, force kill")
		}()
		// Run
		runner, err := runner.CreateRunner(p, args)
		if err != nil {
			logger.Fatal(err)
			return
		}
		go runner.Run(projectDir, stopChan, executedChan, errChan)
		// wait for errChan
		<-executedChan
		err = <-errChan
		if err != nil {
			logger.Fatal(err)
		}
	},
}
