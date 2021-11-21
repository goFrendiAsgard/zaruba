package pathcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "path",
	Short: "path manipulation utilities",
}

func Init() {
	Cmd.AddCommand(getEnvCmd)
	Cmd.AddCommand(getPortConfigCmd)
	Cmd.AddCommand(getAppNameCmd)
	Cmd.AddCommand(getRelativePathCmd)
}
