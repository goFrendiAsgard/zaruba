package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/watcher"
)

func init() {
	rootCmd.AddCommand(watchCmd)
}

var watchCmd = &cobra.Command{
	Use:   "watch [project-dir [...args]]",
	Short: "Watch project and organize accordingly",
	Long:  `Zaruba will perform "organize" whenever something changed`,
	Run: func(cmd *cobra.Command, args []string) {
		// get projectDir
		projectDir := "."
		arguments := []string{}
		if len(args) > 0 {
			projectDir = args[0]
			if len(args) > 1 {
				arguments = args[1:]
			}
		}
		// invoke action
		stopChan := make(chan bool)
		errChan := make(chan error)
		go watcher.Watch(projectDir, stopChan, errChan, arguments...)
		// listen to kill signal
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			stopChan <- true
		}()
		// wait for errChan
		err := <-errChan
		if err != nil {
			log.Fatal("[ERROR] ", err)
		}
	},
}
