package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/cmd/advertisementcmd"
	"github.com/state-alchemists/zaruba/cmd/base64cmd"
	"github.com/state-alchemists/zaruba/cmd/envcmd"
	"github.com/state-alchemists/zaruba/cmd/filecmd"
	"github.com/state-alchemists/zaruba/cmd/installcmd"
	"github.com/state-alchemists/zaruba/cmd/jsoncmd"
	"github.com/state-alchemists/zaruba/cmd/linescmd"
	"github.com/state-alchemists/zaruba/cmd/listcmd"
	"github.com/state-alchemists/zaruba/cmd/mapcmd"
	"github.com/state-alchemists/zaruba/cmd/numcmd"
	"github.com/state-alchemists/zaruba/cmd/pathcmd"
	"github.com/state-alchemists/zaruba/cmd/projectcmd"
	"github.com/state-alchemists/zaruba/cmd/strcmd"
	"github.com/state-alchemists/zaruba/cmd/taskcmd"
	"github.com/state-alchemists/zaruba/cmd/yamlcmd"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var rootLong = fmt.Sprintf(`
MMM"""AMV                               *MM              db      
M'   AMV                                 MM             ;MM:     
'   AMV    ,6"Yb.  '7Mb,od8 '7MM  '7MM   MM,dMMb.      ,V^MM.    
   AMV    8)   MM    MM' "'   MM    MM   MM    'Mb    ,M  'MM    
  AMV   ,  ,pm9MM    MM       MM    MM   MM     M8    AbmmmqMA   
 AMV   ,M 8M   MM    MM       MM    MM   MM.   ,M9   A'     VML  
AMVmmmmMM 'Moo9^Yo..JMML.     'Mbod"YML. P^YbmdP'  .AMA.   .AMMA.
--.. .- .-. ..- -... .-    .--. .-.. . .- ... .    ... - .- .-. - 
                                      Task runner and CLI utility
%s
`, ZarubaVersion)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zaruba",
	Short: "Task runner and CLI utility",
	Long:  rootLong,
}

func init() {
	dsl.SetDefaultEnv()
	rootCmd.AddCommand(tocCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(pleaseCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(advertisementcmd.Cmd)
	advertisementcmd.Init()
	rootCmd.AddCommand(base64cmd.Cmd)
	base64cmd.Init()
	rootCmd.AddCommand(envcmd.Cmd)
	envcmd.Init()
	rootCmd.AddCommand(filecmd.Cmd)
	filecmd.Init()
	rootCmd.AddCommand(installcmd.Cmd)
	installcmd.Init()
	rootCmd.AddCommand(jsoncmd.Cmd)
	jsoncmd.Init()
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
