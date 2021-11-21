package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/cmd/advertisementcmd"
	"github.com/state-alchemists/zaruba/cmd/envcmd"
	"github.com/state-alchemists/zaruba/cmd/installcmd"
	"github.com/state-alchemists/zaruba/cmd/linescmd"
	"github.com/state-alchemists/zaruba/cmd/listcmd"
	"github.com/state-alchemists/zaruba/cmd/mapcmd"
	"github.com/state-alchemists/zaruba/cmd/numcmd"
	"github.com/state-alchemists/zaruba/cmd/pathcmd"
	"github.com/state-alchemists/zaruba/cmd/projectcmd"
	"github.com/state-alchemists/zaruba/cmd/strcmd"
	"github.com/state-alchemists/zaruba/cmd/taskcmd"
	"github.com/state-alchemists/zaruba/cmd/yamlcmd"
	"github.com/state-alchemists/zaruba/output"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zaruba",
	Short: "Task runner and CLI utilities",
	Long: `
 _____                _       _
|__  /__ _ _ __ _   _| |__   / \
  / // _  |  __| | | |  _ \ / _ \
 / /| (_| | |  | |_| | |_) / ___ \
/____\__,_|_|   \__,_|_.__/_/   \_\
Task runner framework and CLI utilities`,
}

func init() {
	executable, _ := os.Executable()
	if os.Getenv("ZARUBA_HOME") == "" {
		os.Setenv("ZARUBA_HOME", filepath.Dir(executable))
	}
	if os.Getenv("ZARUBA_BIN") == "" {
		os.Setenv("ZARUBA_BIN", executable)
	}
	if os.Getenv("ZARUBA_SHELL") == "" {
		os.Setenv("ZARUBA_SHELL", "bash")
	}

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(pleaseCmd)
	rootCmd.AddCommand(serveCmd)

	rootCmd.AddCommand(advertisementcmd.Cmd)
	advertisementcmd.Init()

	rootCmd.AddCommand(envcmd.Cmd)
	envcmd.Init()

	rootCmd.AddCommand(installcmd.Cmd)
	installcmd.Init()

	rootCmd.AddCommand(linescmd.Cmd)
	linescmd.Init()

	rootCmd.AddCommand(listcmd.Cmd)
	listcmd.Init()

	rootCmd.AddCommand(mapcmd.Cmd)
	mapcmd.Init()

	rootCmd.AddCommand(numcmd.Cmd)
	numcmd.Init()

	rootCmd.AddCommand(pathcmd.Cmd)
	pathcmd.Init()

	rootCmd.AddCommand(projectcmd.Cmd)
	projectcmd.Init()

	rootCmd.AddCommand(strcmd.Cmd)
	strcmd.Init()

	rootCmd.AddCommand(taskcmd.Cmd)
	taskcmd.Init()

	rootCmd.AddCommand(yamlcmd.Cmd)
	yamlcmd.Init()

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		logger.Println(err)
		logger.DPrintfError("To run a task you need to type %s%szaruba please <task-name>%s\n", decoration.Bold, decoration.Yellow, decoration.Normal)
		os.Exit(1)
	}
}
