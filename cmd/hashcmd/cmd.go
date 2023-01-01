package hashcmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "hash",
	Short: "hash related utilities",
}

func Init() {
	Cmd.AddCommand(md5Cmd)
	Cmd.AddCommand(sha1Cmd)
	Cmd.AddCommand(sha256Cmd)
}
