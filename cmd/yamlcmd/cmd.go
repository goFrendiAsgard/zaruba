package yamlcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "yaml",
	Short: "Yaml manipulation utilities",
}

func Init() {
	Cmd.AddCommand(readCmd)
	Cmd.AddCommand(printCmd)
}
