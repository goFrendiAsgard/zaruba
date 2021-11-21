package envcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "env",
	Short: "Env utilities",
}

func Init() {
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(readCmd)
	Cmd.AddCommand(printCmd)
	Cmd.AddCommand(writeCmd)
}
