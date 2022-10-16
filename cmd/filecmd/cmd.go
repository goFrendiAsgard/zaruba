package filecmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "file",
	Short: "File utilities",
}

func Init() {
	Cmd.AddCommand(copyCmd)
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(walkCmd)
}
