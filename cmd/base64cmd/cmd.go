package base64cmd

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64 manipulation utilities",
}

func Init() {
	Cmd.AddCommand(decodeCmd)
	Cmd.AddCommand(encodeCmd)
}
